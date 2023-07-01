package main

import (
	"flag"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"sync"
	"time"
	pb "web_hw1/pbGenerated"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "web_hw1/Gateway/server/docs"
)

var (
	addrAuth = flag.String("addrAuth", "localhost:5052", "the address to connect to")
	addrBiz  = flag.String("addrBiz", "localhost:5062", "the address to connect to")
)

type RequestCounter struct {
	Counters map[string][]int64
	Lock     sync.Mutex
}

type BlockedUsers struct {
	information map[string]int64
	Lock        sync.Mutex
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{
		Counters: make(map[string][]int64),
		Lock:     sync.Mutex{},
	}
}

func NewBlockedUsers() *BlockedUsers {
	return &BlockedUsers{
		information: make(map[string]int64),
		Lock:        sync.Mutex{},
	}
}

func (requestCounter *RequestCounter) IncrementIP(ip string) {
	requestCounter.Lock.Lock()
	defer requestCounter.Lock.Unlock()
	now := time.Now().Unix()
	requestCounter.Counters[ip] = append(requestCounter.Counters[ip], now)
	// cleaning up expired entries
	cleanupThreshold := now - 1
	for i := 0; i < len(requestCounter.Counters[ip]); i++ {
		if requestCounter.Counters[ip][i] < cleanupThreshold {
			requestCounter.Counters[ip] = requestCounter.Counters[ip][i+1:]
			i--
		}
	}
}

func (blockedUsers *BlockedUsers) blockIP(ip string) {
	blockedUsers.Lock.Lock()
	defer blockedUsers.Lock.Unlock()
	now := time.Now().Unix()
	blockedUsers.information[ip] = now
}

// IsRateLimited : find if the user has requested more than 100 times in a second
func (requestCounter *RequestCounter) IsRateLimited(ip string) bool {
	requestCounter.Lock.Lock()
	defer requestCounter.Lock.Unlock()
	counter := len(requestCounter.Counters[ip])
	if counter > 100 {
		return true
	}
	return false
}

func (blockedUsers *BlockedUsers) IsBlocked(ip string) bool {
	blockedUsers.Lock.Lock()
	defer blockedUsers.Lock.Unlock()
	value, ok := blockedUsers.information[ip]
	if ok {
		now := time.Now().Unix()
		if (now - value) > 86400 {
			// Unblock the ip
			delete(blockedUsers.information, ip)
			return false
		}
		return true
	}
	return false
}

func (requestCounter *RequestCounter) Middleware(blockedUsers *BlockedUsers) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// increment the request counter
		requestCounter.IncrementIP(ip)
		// check if the user's limit has exceeded
		if requestCounter.IsRateLimited(ip) {
			blockedUsers.blockIP(ip)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Blocked for 24 hours.",
			})
			return
		}
		// continue to the next middleware handler
		c.Next()
	}
}

func (blockedUsers *BlockedUsers) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if blockedUsers.IsBlocked(ip) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "You are blocked because of too many requests.",
			})
			return
		}
		c.Next()
	}
}

// @Summary GetPq
// @Description Get Pq from server
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body pb.ReqPq_Request true "Request body"
// @Success 200 {object} pb.ReqPq_Response
// @Failure 400 {object} map[string]interface{}
// @Router /auth/req_pq [post]
func getPq(c *gin.Context) {
	flag.Parse()
	var req pb.ReqPq_Request
	if err := c.BindJSON(&req); err != nil {
		return
	}

	conn, err := grpc.Dial(*addrAuth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewAuth_ServiceClient(conn)

	res, err := client.ReqPq(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.ReqPq_Response{
		Nonce:       res.GetNonce(),
		ServerNonce: res.GetServerNonce(),
		MessageId:   res.GetMessageId(),
		P:           res.GetP(),
		G:           res.GetG(),
	}
	c.IndentedJSON(http.StatusOK, response)
}

// @Summary GetDh
// @Description Get Dq from server
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body pb.ReqDh_Request true "Request body"
// @Success 200 {object} pb.ReqDh_Response
// @Failure 400 {object} map[string]interface{}
// @Router /auth/req_dh [post]
func getDh(c *gin.Context) {

	flag.Parse()
	var req pb.ReqDh_Request
	if err := c.BindJSON(&req); err != nil {
		return
	}
	fmt.Println("req server nonce is " + req.ServerNonce)
	conn, err := grpc.Dial(*addrAuth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewAuth_ServiceClient(conn)

	res, err := client.Req_DHParams(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.ReqDh_Response{
		Nonce:       res.GetNonce(),
		ServerNonce: res.GetServerNonce(),
		MessageId:   res.GetMessageId(),
		B:           res.GetB(),
	}
	c.IndentedJSON(http.StatusOK, response)
}

// @Summary GetUser
// @Description Get User from server
// @Tags Biz
// @Accept json
// @Produce json
// @Param request body pb.GetUsersRequest true "Request body"
// @Success 200 {object} pb.GetUsersResponse
// @Failure 400 {object} map[string]interface{}
// @Router /get-user [post]
func getUser(c *gin.Context) {
	flag.Parse()
	var req pb.GetUsersRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	// validate the authKey
	request := pb.Seek{
		AuthKey: req.GetAuthKey(),
	}

	conn, err := grpc.Dial(*addrAuth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewAuth_ServiceClient(conn)
	res, err := client.Authenticate(c, &request)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !res.GetValid() {
		log.Print("authentication failed")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "authentication failed",
		})
		return
	}

	BizConn, err := grpc.Dial(*addrBiz, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer BizConn.Close()
	BizClient := pb.NewGetUsersServiceClient(BizConn)

	resp, err := BizClient.GetUsers(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.GetUsersResponse{
		Users:     resp.Users,
		MessageId: 0,
	}
	c.IndentedJSON(http.StatusOK, response)
}

// @Summary GetUserWithInjection
// @Description Get User With Injection from server
// @Tags Biz
// @Accept json
// @Produce json
// @Param request body pb.GetUsersWithInjectionRequest true "Request body"
// @Success 200 {object} pb.GetUsersResponse
// @Failure 400 {object} map[string]interface{}
// @Router /get-user-inject [post]
func getUserInjection(c *gin.Context) {
	flag.Parse()
	var req pb.GetUsersWithInjectionRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}

	request := pb.Seek{
		AuthKey: req.GetAuthKey(),
	}

	conn, err := grpc.Dial(*addrAuth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewAuth_ServiceClient(conn)
	res, err := client.Authenticate(c, &request)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !res.GetValid() {
		log.Print("authentication failed")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "authentication failed",
		})
		return
	}

	bizConn, err := grpc.Dial(*addrBiz, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer bizConn.Close()
	bizClient := pb.NewGetUsersServiceClient(bizConn)

	resp, err := bizClient.GetUsersWithInjection(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.GetUsersResponse{
		Users:     resp.Users,
		MessageId: 0,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	// middleware to handle too many request
	counter := NewRequestCounter()
	blockedUsers := NewBlockedUsers()
	router.Use(counter.Middleware(blockedUsers))
	router.Use(blockedUsers.Middleware())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/auth/req_pq", getPq)
	router.POST("/auth/req_dh", getDh)
	router.POST("/get-user", getUser)
	router.POST("/get-user-inject", getUserInjection)
	router.Run("localhost:6433")
}
