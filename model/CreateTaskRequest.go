package model

type CreateTaskRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Priority int    `json:"priority"`
}
