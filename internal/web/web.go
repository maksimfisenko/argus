package web

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/maksimfisenko/argus/internal/db"
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

func Start(database *sql.DB) {
	tmpl := template.Must(template.ParseFiles("./internal/web/template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		snaps, err := db.Fetch(database)
		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, snaps)
	})

	log.Println("Web server started on 8080")
	http.ListenAndServe(":8080", nil)
}
