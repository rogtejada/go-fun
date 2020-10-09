package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Received request:", r.Method, r.URL)

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}
