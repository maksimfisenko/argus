package agent

import (
	"context"

	"github.com/maksimfisenko/argus/internal/metrics"
	"github.com/maksimfisenko/argus/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Sender struct {
	conn   *grpc.ClientConn
	client proto.ArgusServiceClient
}

func NewSender(address string) (*Sender, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewArgusServiceClient(conn)

	return &Sender{conn: conn, client: client}, nil
}

func (s *Sender) Close() {
	s.conn.Close()
}

func (s *Sender) SendSnaphot(ctx context.Context, snap metrics.Snapshot) error {
	_, err := s.client.SendSnapshot(ctx, &proto.Snapshot{
		Cpu:    snap.CPU,
		Memory: snap.Memory,
	})
	return err
}
