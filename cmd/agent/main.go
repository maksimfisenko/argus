package main

import (
	"context"
	"os"
	"time"

	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/metrics"
	"github.com/maksimfisenko/argus/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Warnf("invalid log level '%s', defaulting to 'info'", cfg.LogLevel)
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	if cfg.LogFile != "stdout" {
		file, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			logrus.Fatalf("cannot open log file: %v", err)
		}
		logrus.SetOutput(file)
		defer file.Close()
	}

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := proto.NewArgusServiceClient(conn)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	logrus.Infof("agent started with interval %s, sending metrics to ':50051'", cfg.Interval)

	for {
		<-ticker.C

		snap, err := metrics.Collect()
		if err != nil {
			logrus.WithError(err).Error("failed to collect metrics")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		resp, err := client.SendSnapshot(ctx, &proto.Snapshot{
			Cpu:    snap.CPU,
			Memory: snap.Memory,
		})

		cancel()

		if err != nil {
			logrus.WithError(err).Error("failed to send snapshot")
			continue
		}

		logrus.Infof("metrics sent successfully, server response: %s", resp.Message)
	}
}
