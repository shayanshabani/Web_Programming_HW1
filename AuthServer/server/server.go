package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	pb "web_hw1/pbGenerated"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

const (
	port = ":5052"
)

type Server struct {
	pb.UnimplementedAuth_ServiceServer
}

// connect to redis server
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var sessionTTL = time.Minute * 20

func generateSHA1Key(input string) string {
	// create new SHA hash object
	sha1Hash := sha1.New()
	// convert the input string to bytes and write it to the hash object
	sha1Hash.Write([]byte(input))
	// get the sha1 hash sum
	hashSum := sha1Hash.Sum(nil)
	// convert the hash sum to a hex string
	sha1Key := hex.EncodeToString(hashSum)
	return sha1Key
}

// generate a random string with specified length
func generate_random_string(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	array := make([]byte, length)
	for i := range array {
		array[i] = charset[rand.Intn(len(charset))]
	}
	return string(array)
}

// generage a random positive odd integer
func generate_random_positive_odd_integer() int32 {
	return rand.Int31n(100)*2 + 1
}

// generate a random prime number
func generate_random_prime_number() int32 {
	for {
		number := rand.Int31n(100) + 1
		if big.NewInt(int64(number)).ProbablyPrime(0) {
			return number
		}
	}
}

func (server *Server) ReqPq(ctx context.Context, request *pb.ReqPq_Request) (*pb.ReqPq_Response, error) {
	// getting inputs:
	client_random_string := request.Nonce
	//client_message_id := request.message_id

	// computing the outputs:

	// generating a random string of size 20
	random_server_side_string := generate_random_string(20)
	// generating a random positive odd number for response's message_id
	random_server_side_integer := generate_random_positive_odd_integer()
	// generating a random prime number like 23
	random_server_side_prime := generate_random_prime_number()
	// generating a random generating number like 5
	random_server_side_generating := rand.Int31n(100)
	// wrapping the response arguments
	response := &pb.ReqPq_Response{
		Nonce:       client_random_string,
		ServerNonce: random_server_side_string,
		MessageId:   random_server_side_integer,
		P:           random_server_side_prime,
		G:           random_server_side_generating,
	}

	// first nonce then server_nonce
	concatenated := string(request.Nonce) + random_server_side_string
	sha1Key := generateSHA1Key(concatenated)
	client_id := strconv.Itoa(int(request.MessageId))
	server_id := strconv.Itoa(int(random_server_side_integer))
	server_prime := strconv.Itoa(int(random_server_side_prime))
	server_generator := strconv.Itoa(int(random_server_side_generating))
	session := map[string]string{
		"nonce":             request.Nonce,
		"message_id":        client_id,
		"server_nonce":      random_server_side_string,
		"server_message_id": server_id,
		"p":                 server_prime,
		"g":                 server_generator}
	for key, value := range session {
		err := client.HSet(context.Background(), sha1Key, key, value).Err()
		if err != nil {
			fmt.Println("couldn't save the map of string to string in redis")
			return nil, fmt.Errorf("saving to redis failed %v", err)
		}
	}

	// set expiration time for the key
	err := client.Expire(context.Background(), sha1Key, sessionTTL).Err()
	if err != nil {
		fmt.Println("couldn't set the expiration time as 20 minutes")
		return nil, fmt.Errorf("set expiration failed %v", err)
	}

	return response, nil
}

func (server *Server) Req_DHParams(ctx context.Context, request *pb.ReqDh_Request) (*pb.ReqDh_Response, error) {
	// getting inputs:

	nonce := request.Nonce
	server_nonce := request.ServerNonce
	A := request.A

	// find the key to get what server saved for this client in redis
	concatenated := nonce + server_nonce
	fmt.Println("server nonce is : " + server_nonce)
	sha1Key := generateSHA1Key(concatenated)

	// get the map of string to string that server saved
	userSession := client.HGetAll(context.Background(), sha1Key).Val()
	if len(userSession) == 0 {
		fmt.Println("wrong key entered")
		return nil, fmt.Errorf("wrong key entered")
	}
	fmt.Println(sha1Key)
	fmt.Println(userSession)

	// extract p & g from the map
	string_p := userSession["p"]
	string_g := userSession["g"]

	raw_p, err := strconv.Atoi(string_p)
	if err != nil {
		fmt.Println("p is not integer in redis")
		return nil, fmt.Errorf("p is not integer")
	}
	p := int32(raw_p)

	raw_g, err := strconv.Atoi(string_g)
	if err != nil {
		fmt.Println("g is not integer in redis")
		return nil, fmt.Errorf("g is not integer")
	}
	g := int32(raw_g)
	// a private key
	b := rand.Int31n(100)

	gBig := big.NewInt(int64(g))
	bBig := big.NewInt(int64(b))
	pBig := big.NewInt(int64(p))

	// Calculate (g^b) % p
	resultBig := new(big.Int).Exp(gBig, bBig, pBig)

	// Convert the result back to int32
	result := int32(resultBig.Int64())

	// wrapping the response
	response := &pb.ReqDh_Response{
		Nonce:       nonce,
		ServerNonce: server_nonce,
		MessageId:   generate_random_positive_odd_integer(),
		B:           result,
	}
	// delete the previous information of the user in the redis
	err = client.Del(context.Background(), sha1Key).Err()
	if err != nil {
		fmt.Println("couldn't delete previous information from redis")
		return nil, fmt.Errorf("delete information from redis failed")
	}

	// calculate common key using the formula key = A^b mod p
	ABig := big.NewInt(int64(A))
	another_resultBig := new(big.Int).Exp(ABig, bBig, pBig)
	common_key := int32(another_resultBig.Int64())
	authKey := strconv.Itoa(int(common_key))
	// store the common key in redis
	err = client.Set(context.Background(), authKey, 1, 0).Err()
	if err != nil {
		fmt.Println("couldn't save the common key on redis")
		return nil, fmt.Errorf("save on redis failed")
	}

	return response, nil
}

func (server *Server) Authenticate(ctx context.Context, request *pb.Seek) (*pb.Confirm, error) {
	// get the authKey from request and make it a string
	authKey := strconv.Itoa(int(request.GetAuthKey()))
	// check if this key exists in redis, in this case exists would be 1, o.w it would be 0
	exists, err := client.Exists(ctx, authKey).Result()
	if err != nil {
		fmt.Println("couldn't read from the redis")
		return nil, fmt.Errorf("read from redis failed")
	}
	yes := false
	// if the key exists
	if exists == 1 {
		yes = true
	}
	response := &pb.Confirm{
		Valid: yes,
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAuth_ServiceServer(server, &Server{})
	fmt.Printf("Server listening on port %s\n", port)

	wait_group := sync.WaitGroup{}
	wait_group.Add(1)

	// start a goroutine for serving each client connection
	go func() {
		defer wait_group.Done()
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	wait_group.Wait()
}
