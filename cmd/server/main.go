package main

import (
	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/logger"
	"github.com/maksimfisenko/argus/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	var cfg config.Server
	if err := config.Load("cmd/server/config.yaml", &cfg); err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	if err := logger.Init(cfg.LogLevel); err != nil {
		logrus.Fatalf("Failed to init logger: %v", err)
	}

	if err := server.Run(&cfg); err != nil {
		logrus.Fatalf("Server exited with error: %v", err)
	}
}
