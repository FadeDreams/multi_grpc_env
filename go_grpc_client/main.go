package main

import (
	"context"
	"log"
	"os"

	pb "github.com/fadedreams/multi_grpc_env/proto" // Import your generated package
	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	//conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "go-grpc-server:8080" // Fallback
	}
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewUserServiceClient(conn)

	// Create a request for CreateUser
	createUserReq := &pb.CreateUserRequest{
		Name:     "Fadedreams",
		Location: "Nakoja",
		Title:    "Software Engineer",
	}

	// Call the CreateUser RPC method
	createUserResp, err := client.CreateUser(context.Background(), createUserReq)
	if err != nil {
		log.Fatalf("Error calling CreateUser: %v", err)
	}

	// Handle the response
	log.Printf("CreateUser Response: %s", createUserResp.Data)

	// Call the GetAllUsers RPC method
	getAllUsersReq := &pb.Empty{}
	getAllUsersResp, err := client.GetAllUsers(context.Background(), getAllUsersReq)
	if err != nil {
		log.Fatalf("Error calling GetAllUsers: %v", err)
	}

	// Handle the response
	log.Printf("All Users:")
	for _, user := range getAllUsersResp.Users {
		log.Printf("User ID: %s, Name: %s, Location: %s, Title: %s\n", user.Id, user.Name, user.Location, user.Title)
	}
}
