package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGRPCServerOnAddr(address string, configureServer func(server *grpc.Server)) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	configureServer(server)

	log.Printf("Starting gRPC server on %s...\n", address)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
