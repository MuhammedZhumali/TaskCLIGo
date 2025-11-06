package model

type CreateTaskRequest struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Priority int    `json:"priority"`
}
