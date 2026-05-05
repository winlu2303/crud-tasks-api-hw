package models

import "time"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at,omitempty"`
}

func NewTask(title string) Task {
	return Task{
		Title:     title,
		Done:      false,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
