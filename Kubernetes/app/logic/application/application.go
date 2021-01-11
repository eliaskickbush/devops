package application

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"kubertest/logic/application/rates"
	"kubertest/logic/application/rates/cache"
	"kubertest/logic/application/rates/openexchange"
	"log"
	"net/http"
	"net/url"
	"os"
)



func Register(r *gin.Engine, redisClient *redis.Client, appID string, baseURL url.URL){

	logger := log.New(os.Stdout, "exchange_rate", log.Ltime | log.Ldate)

	cacheDao := cache.Dao{
		RedisClient: redisClient,
		Logger:      logger,
	}

	openExchangeDao := openexchange.Dao{
		AppID:      appID,
		BaseURL:    baseURL,
		Logger:     logger,
		HTTPClient: http.Client{},
	}

	openExchangeService := openexchange.Service{
		Dao:    openExchangeDao,
		Logger: logger,
	}

	service := rates.Service{
		Base:                "",
		CacheDao:            cacheDao,
		OpenExchangeService: openExchangeService,
		Logger:              logger,
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/update-cache", BuildUpdateCacheHandler(service))
		v1.GET("/convert", BuildConvertHandler(service))
	}
}