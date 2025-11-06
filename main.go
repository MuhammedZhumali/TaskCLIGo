package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"task-manager/controller"
	"task-manager/repo"
	"task-manager/service"
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

	// Initialize layers
	taskRepo := repo.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService)

	// Setup router
	router := gin.Default()

	// Routes
	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks", taskController.GetAllTasks)
	router.GET("/tasks/:id", taskController.GetTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}
