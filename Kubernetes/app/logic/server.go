package logic

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"kubertest/logic/application"
	"kubertest/logic/healthcheck"
	"net/url"
	"os"
	"strconv"
)

type RedisConfig struct {
	Host string
	Port int
}

func GetRedisConfig() (*RedisConfig, error){
	cfg := &RedisConfig{}
	cfg.Host = os.Getenv("REDIS_ADDRESS")
	if len(cfg.Host) == 0 {
		return nil,errors.New("error fetching redis address")
	}
	strPort := os.Getenv("REDIS_PORT")
	res, err := strconv.Atoi(strPort)
	if err != nil {
		return nil, errors.New("could not convert given port to int")
	}
	cfg.Port = res
	if len(cfg.Host) == 0 {
		return nil,errors.New("error fetching redis port")
	}
	return cfg,nil
}

// StartApp ...
func StartApp() {
	router := gin.Default()

	redisCfg, err := GetRedisConfig()
	if err != nil {
		panic(err.Error())
	}
	redisClient,err := GetRedisConnection(redisCfg.Host, redisCfg.Port)
	if err != nil {
		panic(err.Error())
	}

	appID,err := GetAppID()
	if err != nil {
		panic(err.Error())
	}
	baseURL,err := GetBaseURL()
	if err != nil {
		panic(err.Error())
	}

	healthcheck.Register(router)
	application.Register(router, redisClient, appID, *baseURL)

	router.Run(":8080")
}

func GetRedisConnection(host string, port int) (*redis.Client,error){
	address := fmt.Sprintf("%s:%d", host,port)
	fmt.Printf("%s\n", address)
	options := &redis.Options{
		Addr: address,
	}
	client := redis.NewClient(options)
	cmd := client.Ping()
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return client, nil
}

func GetAppID() (string,error){
	appID := os.Getenv("APP_ID")
	if len(appID) == 0 {
		return "", errors.New("empty app id variable")
	}
	return appID, nil
}

func GetBaseURL() (*url.URL, error) {
	baseURL,err := url.Parse(os.Getenv("BASE_URL"))
	if err != nil {
		return nil,errors.New("there was an error building base URL")
	}
	return baseURL, nil
}