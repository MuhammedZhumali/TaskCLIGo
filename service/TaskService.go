package service

import "task-manager/model"

type TaskService interface {
	CreateTask(task model.CreateTaskRequest) (
		model.TaskResponse,
		error,
	)

	UpdateTask(id int, request model.UpdateTaskRequest) (
		model.TaskResponse,
		error,
	)

	GetTaskById(id int) (
		model.TaskResponse,
		error,
	)

	GetAllTasks() (
		[]model.TaskResponse,
		error,
	)

	DeleteTask(id int) error
}
