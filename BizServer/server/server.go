package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"web_hw1/BizServer/database"
	pb "web_hw1/pbGenerated"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5062, "The server port")
)

type server struct {
	pb.UnimplementedGetUsersServiceServer
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
}

func (s *server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	userId := req.UserId
	if userId == 0 {
		users, err := database.GetFirst100UsersFromDatabase(s.dbPool)
		if err != nil {
			return nil, err
		}
		return &pb.GetUsersResponse{Users: users, MessageId: rand.Int31n(100)*2 + 1}, nil
	}

	users, err := database.GetUsersFromDatabase(s.dbPool, userId)
	if err != nil {
		return nil, err
	}
	return &pb.GetUsersResponse{Users: users, MessageId: rand.Int31n(100)*2 + 1}, nil
}

func (s *server) GetUsersWithInjection(ctx context.Context, req *pb.GetUsersWithInjectionRequest) (*pb.GetUsersResponse, error) {
	userId := req.UserId

	if userId == "" {
		users, err := database.GetFirst100UsersFromDatabase(s.dbPool)
		if err != nil {
			return nil, err
		}
		return &pb.GetUsersResponse{Users: users, MessageId: rand.Int31n(100)*2 + 1}, nil
	} else {
		users, err := database.GetUsersWithInjectionFromDatabase(s.dbPool, userId)
		if err != nil {
			return nil, err
		}
		return &pb.GetUsersResponse{Users: users, MessageId: rand.Int31n(100)*2 + 1}, nil
	}
}

func main() {
	flag.Parse()
	pool, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetUsersServiceServer(s, &server{dbPool: pool, redisClient: client})

	wait_group := sync.WaitGroup{}
	wait_group.Add(1)
	go func() {
		defer wait_group.Done()
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	wait_group.Wait()
}
