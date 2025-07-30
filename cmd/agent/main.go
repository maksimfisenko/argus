package main

import (
	"github.com/maksimfisenko/argus/internal/agent"
	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	if err := logger.Init(cfg.LogLevel, cfg.LogFile); err != nil {
		logrus.Fatalf("failed to init logger: %v", err)
	}

	if err := agent.Run(cfg); err != nil {
		logrus.Fatalf("agent exited with error: %v", err)
	}
}
