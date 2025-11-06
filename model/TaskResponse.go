package model

import "time"

type TaskResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Owner      string    `json:"owner"`
	Priority   int       `json:"priority"`
	Created_at time.Time `json:"created_at"`
}
