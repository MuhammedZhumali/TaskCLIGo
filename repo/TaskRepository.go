package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"task-manager/model"
	"time"
)

type TaskRepository interface {
	CreateTask(task model.CreateTaskRequest) (model.TaskResponse, error)
	UpdateTask(task model.Task) (model.Task, error)
	GetTaskById(id int) (model.Task, error)
	GetAllTasks() ([]model.Task, error)
	DeleteTask(id int) (string, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task model.CreateTaskRequest) (model.TaskResponse, error) {
	query := `
        INSERT INTO tasks (name, owner, priority)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `

	var id int
	var createdAt time.Time
	err := r.db.QueryRow(query, task.Name, task.Owner, task.Priority).Scan(&id, &createdAt)

	if err != nil {
		return model.TaskResponse{}, err
	}

	return model.TaskResponse{
		ID:         id,
		Name:       task.Name,
		Owner:      task.Owner,
		Priority:   task.Priority,
		Created_at: createdAt,
	}, nil
}

func (r *taskRepository) GetTaskById(id int) (model.Task, error) {
	query := `
        SELECT id, name, owner, priority, created_at
        FROM tasks
        WHERE id = $1
    `

	var task model.Task

	err := r.db.QueryRow(query, id).
		Scan(&task.ID, &task.Name, &task.Owner, &task.Priority, &task.Created_at)

	if err == sql.ErrNoRows {
		return model.Task{}, errors.New("task not found")
	}

	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]model.Task, error) {
	query := `
	SELECT id, name, owner, priority, created_at
	FROM tasks ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task

		err := rows.Scan(&task.ID, &task.Name, &task.Owner, &task.Priority, &task.Created_at)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) UpdateTask(task model.Task) (model.Task, error) {
	query := `UPDATE tasks SET name = $1, owner = $2, priority = $3
	WHERE id = $4
	RETURNING id, name, owner, priority, created_at
	`
	var updatedTask model.Task

	err := r.db.QueryRow(query, task.Name, task.Owner, task.Priority, task.ID).
		Scan(&updatedTask.ID, &updatedTask.Name, &updatedTask.Owner, &updatedTask.Priority, &updatedTask.Created_at)

	if err == sql.ErrNoRows {
		return model.Task{}, errors.New("task not found")
	}

	if err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepository) DeleteTask(id int) (msg string, err error) {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return "Error occurred", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "Error occurred", err
	}

	if rowsAffected == 0 {
		return "Task not found", errors.New("task not found")
	}

	return "Task deleted successfully by id:" + fmt.Sprintf("%d", id), nil
}
