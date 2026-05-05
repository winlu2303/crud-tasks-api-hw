package main

import (
	"log"
	"net/http"

	"crud-tasks-api/internal/handlers"
	"crud-tasks-api/internal/middleware"
	"crud-tasks-api/internal/storage"
)

func main() {
	// create storage
	store := storage.NewMemoryStorage()

	// create handler
	h := handlers.New(store)

	// setup routes with middleware
	mux := http.NewServeMux()

	// apply logging middleware
	mux.HandleFunc("/tasks", middleware.Logger(
		middleware.JSONContentType(h.TasksCollection),
	))

	mux.HandleFunc("/tasks/", middleware.Logger(
		middleware.JSONContentType(h.TaskItem),
	))

	mux.HandleFunc("/health", middleware.Logger(
		middleware.JSONContentType(h.HealthCheck),
	))

	log.Println("Server listening on :8080")
	log.Println("Available endpoints:")
	log.Println("	GET		/tasks			- List all tasks")
	log.Println("	POST		/tasks			- Create a new task")
	log.Println("	GET		/tasks/{id}		- Get task by ID")
	log.Println("	PUT		/tasks/{id}		- Update task by ID")
	log.Println("	Delete		/tasks/{id}		- Delete task by ID")
	log.Println("	GET		/health			- Health check")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("X Server failed to start:", err)
	}
}
