package main

func CreateTables() {

	DB.Exec(`CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	)`)

	DB.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		status TEXT DEFAULT 'todo',
		project_id INTEGER
	)`)
}
