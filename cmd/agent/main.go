package main

import (
	"github.com/maksimfisenko/argus/internal/agent"
	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	var cfg config.Agent
	if err := config.Load("cmd/agent/config.yaml", &cfg); err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	if err := logger.Init(cfg.LogLevel); err != nil {
		logrus.Fatalf("Failed to init logger: %v", err)
	}

	if err := agent.Run(&cfg); err != nil {
		logrus.Fatalf("Agent exited with error: %v", err)
	}
}
