CREATE TABLE IF NOT EXISTS snapshots (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    agent_id TEXT,
    cpu REAL,
    memory REAL,
    disk_usage REAL,
    avg_load REAL,
    uptime INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);