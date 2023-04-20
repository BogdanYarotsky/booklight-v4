package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Response struct {
	Query string `json:"query"`
}

func main() {
	endpoint := "localhost:50051"
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(endpoint, creds)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	http.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
				query := r.URL.Query().Get("q")

		if query == "" {
			http.Error(w, "missing query parameter 'q'", http.StatusBadRequest)
			return
		}

		// results := searchService.Search(query)

		response := Response{Query: query}
		// add search results to the response
		// ...
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonBytes))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
