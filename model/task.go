package model

import "time"

type Priority string

const (
	Low    Priority = "Low"
	Medium Priority = "Medium"
	High   Priority = "High"
)

type Task struct {
	id          int
	title       string
	description string
	dueDate     time.Time
	priority    Priority
	createdAt   time.Time
}
