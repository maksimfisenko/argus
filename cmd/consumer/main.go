package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/maksimfisenko/argus/proto"
	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "snapshots",
		GroupID: "argus-consumer-group",
	})

	log.Println("consumer started...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		var snap proto.Snapshot
		if err := json.Unmarshal(msg.Value, &snap); err != nil {
			log.Printf("error unmarshaling snapshot: %v", err)
			continue
		}

		log.Printf("received snapshot from %s: CPU=%.2f, Memory=%.2f", snap.AgentId, snap.Cpu, snap.Memory)
	}
}
