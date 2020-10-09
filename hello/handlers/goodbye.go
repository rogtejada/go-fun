package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Received request:", r.Method, r.URL)
	fmt.Fprintf(rw, "Goodbye")
}
