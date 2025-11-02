package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"task-manager/model"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://taskmanager:taskmanagerpass@localhost:5432/taskmanagerdb?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	t := model.Task{
		Name:     "first task",
		Owner:    "opir",
		Priority: 1,
	}

	query := `
		INSERT INTO tasks (name, owner, priority)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	if err := db.QueryRow(query, t.Name, t.Owner, t.Priority).Scan(&t.ID); err != nil {
		log.Fatal("Failed to insert task:", err)
	}

	fmt.Println("Task inserted successfully, id:", t.ID)

	getAll := `SELECT id, name, owner, priority FROM tasks`

	rows, err := db.Query(getAll)
	if err != nil {
		log.Fatal("Failed to query tasks:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tt model.Task
		if err := rows.Scan(&tt.ID, &tt.Name, &tt.Owner, &tt.Priority); err != nil {
			log.Fatal("Failed to scan task:", err)
		}
		fmt.Printf("%+v\n", tt)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}
}
