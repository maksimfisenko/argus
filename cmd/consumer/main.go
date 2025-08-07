package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/maksimfisenko/argus/proto"
	"github.com/segmentio/kafka-go"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data/argus.db")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

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

		log.Printf("%s: CPU=%.2f, MEM=%.2f", snap.AgentId, snap.Cpu, snap.Memory)

		_, err = db.Exec(`
			INSERT INTO snapshots(agent_id, cpu, memory, disk_usage, avg_load, uptime)
			VALUES (?, ?, ?, ?, ?, ?)`,
			snap.AgentId, snap.Cpu, snap.Memory, snap.DiskUsage, snap.AvgLoad, snap.Uptime,
		)
		if err != nil {
			log.Printf("error inserting to db: %v", err)
		}

		log.Print("successfully inserted")
	}
}
