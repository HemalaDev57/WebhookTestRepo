package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // External dependency
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Dockerized Go World!")
	})
	r.Run(":8080")
}
