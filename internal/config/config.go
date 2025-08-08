package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Agent struct {
	ID            string `yaml:"id"`
	LogLevel      string `yaml:"log_level"`
	ServerAddress string `yaml:"server_address"`
	Interval      int    `yaml:"interval"`
}

type Server struct {
	Address      string   `yaml:"address"`
	LogLevel     string   `yaml:"log_level"`
	KafkaBrokers []string `yaml:"kafka_brokers"`
	KafkaTopic   string   `yaml:"kafka_topic"`
}

type Consumer struct {
	LogLevel     string   `yaml:"log_level"`
	KafkaBrokers []string `yaml:"kafka_brokers"`
	KafkaTopic   string   `yaml:"kafka_topic"`
	KafkaGroupID string   `yaml:"kafka_group_id"`
	DbPath       string   `yaml:"db_path"`
}

func Load(path string, out any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
