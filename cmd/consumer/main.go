package main

import (
	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/consumer"
	"github.com/maksimfisenko/argus/internal/logger"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var cfg config.Consumer
	if err := config.Load("./cmd/consumer/config.yaml", &cfg); err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	if err := logger.Init(cfg.LogLevel); err != nil {
		logrus.Fatalf("Failed to init logger: %v", err)
	}

	if err := consumer.Run(&cfg); err != nil {
		logrus.Fatalf("Consumer exited with error: %v", err)
	}
}
