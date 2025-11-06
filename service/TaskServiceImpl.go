package service

import (
	"errors"
	"task-manager/model"
	"task-manager/repo"
)

type TaskServiceImpl struct {
	repo repo.TaskRepository
}

func NewTaskService(repo repo.TaskRepository) TaskService {
	return &TaskServiceImpl{repo: repo}
}

func (s *TaskServiceImpl) CreateTask(task model.CreateTaskRequest) (model.TaskResponse, error) {
	if task.Name == "" {
		return model.TaskResponse{}, errors.New("task name cannot be empty")
	}
	if task.Owner == "" {
		return model.TaskResponse{}, errors.New("task owner cannot be empty")
	}
	if task.Priority <= 0 {
		return model.TaskResponse{}, errors.New("priority must be greater than 0")
	}

	response, err := s.repo.CreateTask(task)
	if err != nil {
		return model.TaskResponse{}, err
	}
	return response, nil
}

func (s *TaskServiceImpl) UpdateTask(id int, req model.UpdateTaskRequest) (model.TaskResponse, error) {
	if id <= 0 {
		return model.TaskResponse{}, errors.New("invalid task ID")

	}
	if req.Name == "" && req.Owner == "" && req.Priority == 0 {
		return model.TaskResponse{}, errors.New("at least one field must be provided")
	}

	existingTask, err := s.repo.GetTaskById(id)
	if err != nil {
		return model.TaskResponse{}, err
	}

	if req.Name != "" {
		existingTask.Name = req.Name
	}
	if req.Owner != "" {
		existingTask.Owner = req.Owner
	}
	if req.Priority > 0 {
		existingTask.Priority = req.Priority
	}

	updateTask, err := s.repo.UpdateTask(existingTask)
	if err != nil {
		return model.TaskResponse{}, err
	}

	return model.TaskResponse{
		ID:         updateTask.ID,
		Name:       updateTask.Name,
		Owner:      updateTask.Owner,
		Priority:   updateTask.Priority,
		Created_at: updateTask.Created_at,
	}, nil
}

func (s *TaskServiceImpl) GetAllTasks() ([]model.TaskResponse, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	responses := make([]model.TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = model.TaskResponse{
			ID:         task.ID,
			Name:       task.Name,
			Owner:      task.Owner,
			Priority:   task.Priority,
			Created_at: task.Created_at,
		}
	}

	return responses, nil
}

func (s *TaskServiceImpl) GetTaskById(id int) (model.TaskResponse, error) {
	if id <= 0 {
		return model.TaskResponse{}, errors.New("invalid task ID")
	}

	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return model.TaskResponse{}, err
	}

	return model.TaskResponse{
		ID:         task.ID,
		Name:       task.Name,
		Owner:      task.Owner,
		Priority:   task.Priority,
		Created_at: task.Created_at,
	}, nil
}

func (s *TaskServiceImpl) DeleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid task ID")
	}

	_, err := s.repo.DeleteTask(id)
	return err // Repository already handles "not found" errors
}
