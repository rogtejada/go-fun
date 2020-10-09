package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloWordHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWordHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL)

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}
