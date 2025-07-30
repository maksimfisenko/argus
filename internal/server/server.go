package server

import (
	"context"
	"log"

	pb "github.com/maksimfisenko/argus/proto"
)

type Server struct {
	pb.UnimplementedArgusServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SendSnapshot(ctx context.Context, req *pb.Snapshot) (*pb.Ack, error) {
	log.Printf("received snapshot from '%s': CPU=%.2f%%, MEM=%.2f%%", req.AgentId, req.Cpu, req.Memory)

	return &pb.Ack{
		Message: "ok",
	}, nil
}
