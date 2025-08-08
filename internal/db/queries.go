package db

const InsertQuery = `INSERT INTO snapshots(agent_id, cpu, memory, disk_usage, avg_load, uptime)
	VALUES (?, ?, ?, ?, ?, ?)`

const FetchQuery = `SELECT agent_id, cpu, memory, disk_usage, avg_load, uptime, created_at
	FROM snapshots
	ORDER BY created_at DESC
	LIMIT 10`
