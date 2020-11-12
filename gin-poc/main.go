package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello-world"})
}

func main() {
	r := gin.Default()
	r.GET("/hello", Hello)
	err := r.Run()

	if err != nil {
		fmt.Println("Error loading gin server")
	}
}
