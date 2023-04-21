package main

import (
	pb "github.com/BogdanYarotsky/booklight-v4/go-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartGrpcClient(endpoint string) (*grpc.ClientConn, pb.BooksClient, error) {
	credentials := insecure.NewCredentials()
	dialOption := grpc.WithTransportCredentials(credentials)
	conn, err := grpc.Dial(endpoint, dialOption)
	client := pb.NewBooksClient(conn)
	return conn, client, err
}