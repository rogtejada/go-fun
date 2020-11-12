package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello-world"})
}

func (s *Service) Check() (bool, error) {
	return true, nil
}
