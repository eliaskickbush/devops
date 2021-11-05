package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	hostname, err := os.Hostname()
	if err != nil {
		panic(fmt.Sprintf("couldn't retrieve hostname: %s", err.Error()))
	}

	app.GET("/greet/:name", func(c *gin.Context) {
		if name := c.Param("name"); name == "" {
			c.String(http.StatusBadRequest, "name not found in request path")
			return
		} else {
			c.String(http.StatusOK, "%s says: hello %s!", hostname, name)
			return
		}
	})

	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	app.Run(":8080")

}
