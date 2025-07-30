package main

import (
	"log"
	"net"

	"github.com/maksimfisenko/argus/internal/server"
	"google.golang.org/grpc"

	pb "github.com/maksimfisenko/argus/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := server.NewServer()
	pb.RegisterArgusServiceServer(grpcServer, srv)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
