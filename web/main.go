package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Snapshot struct {
	AgentID   string
	CPU       float64
	Memory    float64
	DiskUsage float64
	AvgLoad   float64
	Uptime    uint64
	CreatedAt time.Time
}

func main() {
	db, err := sql.Open("sqlite3", "./data/argus.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tmpl := template.Must(template.ParseFiles("./web/template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT agent_id, cpu, memory, disk_usage, avg_load, uptime, created_at
			FROM snapshots
			ORDER BY created_at DESC
			LIMIT 10`)
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
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

		tmpl.Execute(w, snaps)
	})

	log.Println("Web server started on 8080")
	http.ListenAndServe(":8080", nil)
}
