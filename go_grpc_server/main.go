package main

import (
	"log"
	"net"

	pb "github.com/fadedreams/multi_grpc_env/proto"
	"github.com/fadedreams/multi_grpc_env/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//lis, err := net.Listen("tcp", "[::1]:8080")
	//lis, err := net.Listen("tcp", "localhost:8080")
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server is running on port 8080")
	grpcServer := grpc.NewServer()
	service := &services.UserServiceServer{}

	pb.RegisterUserServiceServer(grpcServer, service)

	// Enable gRPC server reflection
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
