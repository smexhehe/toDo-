package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var Tasks []Task

func taskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		jsonData, err := json.Marshal(Tasks)
		if err != nil {
			http.Error(w, "Error marshaling", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	case http.MethodPost:
		var newTask Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Error marshaling", 500)
			return
		}
		newTask.ID = len(Tasks) + 1
		Tasks = append(Tasks, newTask)
		jsonDataAccept, err := json.Marshal(newTask)
		if err != nil {
			http.Error(w, "Error marshaling", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonDataAccept)
	default:
		http.Error(w, "method not allowed", 405)

	}

}

func taskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id not number", 400)
	}

	switch r.Method {
	case http.MethodPut:

	}
}

func main() {

	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/task/", taskByIDHandler)

	fmt.Println("Start server")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
