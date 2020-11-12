package main

import (
	"flag"
	"fmt"
	"gin-consul-poc/service"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	ttl := flag.Duration("ttl", time.Second*15, "Service TTL check duration")
	flag.Parse()

	s, err := service.New(*ttl)

	r := gin.Default()
	r.GET("/hello", s.Hello)
	err = r.Run()

	if err != nil {
		fmt.Println("Error loading gin server")
	}
}

