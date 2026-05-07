package handlers

import (
	"TODO/internal/models"
	"TODO/internal/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var tasksFile = "data/tasks.json"

func SetTasksFile(path string) {
	tasksFile = path
}

type errorResponse struct {
	Error string `json:"error"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, storage.GetTasks())

	case http.MethodPost:
		var newTask models.Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {

			writeJSONError(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(newTask.Title) == "" {
			writeJSONError(w, "title is required", http.StatusBadRequest)
			return
		}
		createdTask := storage.CreateTask(newTask)

		err = storage.SaveTasks(tasksFile)
		if err != nil {
			writeJSONError(w, "failed to save tasks", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusCreated, createdTask)

	default:
		writeJSONError(w, "method not allowed", http.StatusMethodNotAllowed)

	}

}

func TaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "id not number", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		task, ok := storage.GetTaskByID(id)
		if !ok {
			writeJSONError(w, "id not found", http.StatusNotFound)
			return
		}

		writeJSON(w, http.StatusOK, task)
		return

	case http.MethodPut:
		var updatedTask models.Task
		err := json.NewDecoder(r.Body).Decode(&updatedTask)
		if err != nil {
			writeJSONError(w, "Error decoding", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(updatedTask.Title) == "" {
			writeJSONError(w, "title is required", http.StatusBadRequest)
			return
		}

		task, ok := storage.UpdateTask(id, updatedTask)
		if !ok {
			writeJSONError(w, "task not found", http.StatusNotFound)
			return
		}
		err = storage.SaveTasks(tasksFile)
		if err != nil {
			writeJSONError(w, "failed to save tasks", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, task)

	case http.MethodPatch:
		var patch models.TaskPatch
		err := json.NewDecoder(r.Body).Decode(&patch)
		if err != nil {
			writeJSONError(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if patch.Title != nil && strings.TrimSpace(*patch.Title) == "" {
			writeJSONError(w, "title is required", http.StatusBadRequest)
			return
		}

		task, ok := storage.PatchTask(id, patch)
		if !ok {
			writeJSONError(w, "task not found", http.StatusNotFound)
			return
		}

		err = storage.SaveTasks(tasksFile)
		if err != nil {
			writeJSONError(w, "failed to save tasks", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, task)

	case http.MethodDelete:
		ok := storage.DeleteTask(id)
		if !ok {
			writeJSONError(w, "task not found", http.StatusNotFound)
			return
		}
		err = storage.SaveTasks(tasksFile)
		if err != nil {
			writeJSONError(w, "failed to save tasks", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return

	default:

		writeJSONError(w, "method not allowed", http.StatusMethodNotAllowed)

	}
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse{
		Error: message,
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
