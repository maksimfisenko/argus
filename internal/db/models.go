package db

import "time"

type Snapshot struct {
	AgentID   string
	CPU       float64
	Memory    float64
	DiskUsage float64
	AvgLoad   float64
	Uptime    uint64
	CreatedAt time.Time
}
