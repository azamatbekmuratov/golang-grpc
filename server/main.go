package main

import (
	pb "bookshop/server/pb/inventory"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func getSampleBooks() []*pb.Book {
	sampleBooks := []*pb.Book{
		{
			Title:     "Golang Play projects",
			Author:    "Bekmuratov Azamat",
			PageCount: 42,
		},
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R. Tolkien",
			PageCount: 1234,
		},
	}
	return sampleBooks
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("fatal to serve: %v", err)
	}
}
