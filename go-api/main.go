package main

import (
	"log"
	"net/http"
)

func main() {
	conn, client, err := StartGrpcClient("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	server := &BooksServer{client: client}
	log.Fatal(http.ListenAndServe("localhost:8080", server))
}
