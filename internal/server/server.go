package server

import (
	"context"
	"encoding/json"

	"github.com/maksimfisenko/argus/internal/kafka"
	pb "github.com/maksimfisenko/argus/proto"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedArgusServiceServer
	producer *kafka.Producer
}

func NewServer(producer *kafka.Producer) *Server {
	return &Server{producer: producer}
}

func (s *Server) SendSnapshot(ctx context.Context, req *pb.Snapshot) (*pb.Ack, error) {
	data, err := json.Marshal(req)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal snapshot")
		return nil, err
	}

	err = s.producer.Publish(ctx, data)
	if err != nil {
		logrus.WithError(err).Error("failed to publish snapshot to Kafka")
		return nil, err
	}

	logrus.Infof("snapshot from agent %s queued", req.AgentId)
	return &pb.Ack{Message: "snapshot queued"}, nil
}
