package handlers

import (
	"TODO/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var Tasks []models.Task

type errorResponse struct {
	Error string `json:"error"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		jsonData, err := json.Marshal(Tasks)
		if err != nil {

			writeJSONError(w, "Error marshaling", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	case http.MethodPost:
		var newTask models.Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {

			writeJSONError(w, "Error marshaling", http.StatusInternalServerError)
			return
		}
		if strings.TrimSpace(newTask.Title) == "" {
			writeJSONError(w, "title is required", http.StatusBadRequest)
			return
		}
		newTask.ID = len(Tasks) + 1
		Tasks = append(Tasks, newTask)
		jsonDataAccept, err := json.Marshal(newTask)
		if err != nil {
			writeJSONError(w, "Error marshaling", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonDataAccept)
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
		for i := range Tasks {
			if Tasks[i].ID == id {
				jsonData, err := json.Marshal(Tasks[i])
				if err != nil {
					writeJSONError(w, "Error marshaling", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonData)
				return
			}
		}
		writeJSONError(w, "id not found", http.StatusNotFound)
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

		for i := range Tasks {
			if Tasks[i].ID == id {
				updatedTask.ID = id
				Tasks[i] = updatedTask

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(updatedTask)
				return
			}
		}

		writeJSONError(w, "task not found", http.StatusNotFound)

	case http.MethodDelete:
		for i := range Tasks {
			if Tasks[i].ID == id {
				Tasks = append(Tasks[:i], Tasks[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		writeJSONError(w, "task not found", http.StatusNotFound)

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
