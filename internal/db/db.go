package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Open(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}

func Insert(database *sql.DB, agentId string, cpu, memory, diskUsage, avgLoad float64, uptime uint64) error {
	_, err := database.Exec(InsertQuery, agentId, cpu, memory, diskUsage, avgLoad, uptime)
	if err != nil {
		return err
	}
	return nil
}

func Fetch(database *sql.DB) ([]Snapshot, error) {
	rows, err := database.Query(FetchQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snaps []Snapshot

	for rows.Next() {
		var s Snapshot
		err := rows.Scan(&s.AgentID, &s.CPU, &s.Memory, &s.DiskUsage, &s.AvgLoad, &s.Uptime, &s.CreatedAt)
		if err == nil {
			snaps = append(snaps, s)
		}
	}

	return snaps, nil
}
