package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/BogdanYarotsky/booklight-v4/go-api/pb"
)

type BooksServer struct {
	client pb.BooksClient
}

func (s *BooksServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/search":
		s.handleSearch(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (s *BooksServer) handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing query parameter 'q'", http.StatusBadRequest)
		return
	}

	resp, err := s.client.Get(r.Context(), &pb.BooksRequest{Query: query})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonBytes))
}