package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.Run(":8081")
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hello, World!")
}
