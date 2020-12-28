package main

import (
	"log"
	"net"

	pb "proto-test/pkg/api"
	roomServer "proto-test/pkg/room"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	roomServ := roomServer.NewRoomServer()

	pb.RegisterRoomServiceServer(grpcServer, roomServ)

	log.Println("Starting Grpc Service listening on :9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
