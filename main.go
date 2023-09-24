package main

import (
	"fmt"
	"url-shortener/pkg/handlers"
	redisDB "url-shortener/pkg/third_party/store"

	"github.com/gin-gonic/gin"
)

func main() {

	gin := gin.Default()

	redisClient := redisDB.NewRedisConnection()
	r := redisDB.RedisStruct{
		Client: redisClient,
	}

	fmt.Println("starting server")

	r.Create()

	handlers.UrlShortenerHandler(gin)

	gin.Run("localhost:8080")

}
