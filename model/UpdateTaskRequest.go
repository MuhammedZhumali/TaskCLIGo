package model

type UpdateTaskRequest struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Priority int    `json:"priority"`
}
