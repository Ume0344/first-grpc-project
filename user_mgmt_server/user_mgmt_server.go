package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/Ume0344/first-grpc-project/user_mgmt"
	"google.golang.org/grpc"
)

// server port where service is being hosted
const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())

	var userId int32 = int32(rand.Intn(100000))

	return &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   userId,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed listening to server port: %v", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterUserManagementServer(server, &UserManagementServer{})

	log.Printf("Server listening at %v", listen.Addr())
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve at server address: %v", err.Error())
	}
}
