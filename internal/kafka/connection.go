package kafka

import (
	"github.com/maksimfisenko/argus/internal/config"
	"github.com/segmentio/kafka-go"
)

func Ping(cfg *config.Consumer) error {
	conn, err := kafka.Dial("tcp", cfg.KafkaBrokers[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ReadPartitions(cfg.KafkaTopic)
	if err != nil {
		return err
	}

	return nil
}
