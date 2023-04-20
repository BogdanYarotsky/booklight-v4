package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/BogdanYarotsky/booklight-v4/go-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	endpoint := "0.0.0.0:9090"
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(endpoint, creds)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	booksClient := pb.NewBooksClient(conn)
	http.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")

		if query == "" {
			http.Error(w, "missing query parameter 'q'", http.StatusBadRequest)
			return
		}

		fmt.Println("Sending: " + query);
		resp, err := booksClient.Get(r.Context(), &pb.BooksRequest{Query: query})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Received: " + resp.Books[0].Title)

		jsonBytes, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonBytes))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
