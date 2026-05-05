package handlers

import (
	"crud-tasks-api/internal/models"
	"crud-tasks-api/internal/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	Store storage.Storage
}

func New(s storage.Storage) *Handler {
	return &Handler{Store: s}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

// that to get all tasks
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := h.Store.List()
		writeJSON(w, http.StatusOK, tasks)

	case http.MethodPost:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON format")
			return
		}

		if task.Title == "" {
			writeError(w, http.StatusBadRequest, "Title is required")
			return
		}

		newTask := models.NewTask(task.Title)
		created, err := h.Store.Create(newTask)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		writeJSON(w, http.StatusCreated, created)

	default:
		writeError(w, http.StatusMethodNotAllowed,
			"Method not allowed")
	}
}

// task/{id} like GET, PUT, DELETE
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	// extract ID from path
	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(path)
	if err != nil {
		writeError(w, http.StatusBadRequest,
			"Invalid task ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, exists := h.Store.Get(id)
		if !exists {
			writeError(w, http.StatusNotFound,
				"Task not found")
			return
		}
		writeJSON(w, http.StatusOK, task)

	case http.MethodPut:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest,
				"Invalid JSON format")
			return
		}

		if task.Title == "" {
			writeError(w, http.StatusBadRequest,
				"Title is required")
			return
		}

		updated, err := h.Store.Update(id, task)
		if err != nil {
			if strings.Contains(err.Error(),
				"not found") {
				writeError(w, http.StatusNotFound,
					err.Error())
			} else {
				writeError(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		if err := h.Store.Delete(id); err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}

		writeJSON(w, http.StatusNoContent, nil)

	default:
		writeError(w, http.StatusMethodNotAllowed,
			"Method not allowed")
	}
}

// health check endpoint
func (h *Handler) HealthCheck(w http.ResponseWriter,
	r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed,
			"Method not allowed")
		return
	}
	writeJSON(w, http.StatusOK,
		map[string]string{"status": "ok"})
}
