package server

import (
	"context"
	"encoding/json"
	"errors"
	"net"

	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/kafka"
	pb "github.com/maksimfisenko/argus/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
		logrus.Error("Failed to marshal snapshot")
		return nil, err
	}

	err = s.producer.Publish(ctx, data)
	if err != nil {
		logrus.Error("Failed to publish snapshot to Kafka")
		return nil, nil
	}

	logrus.Infof("Successfully queued snapshot [agent-id='%s']", req.AgentId)

	return &pb.Ack{Message: "OK"}, nil
}

func Run(cfg *config.Server) error {
	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return errors.New("failed to set up new listener")
	}

	grpcServer := grpc.NewServer()

	producer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	defer producer.Close()

	srv := NewServer(producer)
	pb.RegisterArgusServiceServer(grpcServer, srv)

	logrus.Infof("The server is up [adress='%s', kafka-brokers='%v', kafka-topic='%s']", cfg.Address, cfg.KafkaBrokers, cfg.KafkaTopic)

	if err := grpcServer.Serve(lis); err != nil {
		return errors.New("failed to serve")
	}

	return nil
}
