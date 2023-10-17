package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Ume0344/first-grpc-project/user_mgmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Make a connection to server
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client failed to connect: %v", err.Error())
	}
	defer connection.Close()

	client := pb.NewUserManagementClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUsers = make(map[string]int32)
	newUsers["John"] = 26
	newUsers["David"] = 14

	for name, age := range newUsers {
		// Request the server to create a client
		response, err := client.CreateUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Could not create a %s user: %v", name, err.Error())
		}

		log.Printf(`User Details: Name: %s, Age: %d, ID: %d`, response.Name, response.Age, response.Id)
	}
}
