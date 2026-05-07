package main

import (
	"TODO/internal/handlers"
	"TODO/internal/storage"
	"fmt"
	"net/http"
)

func main() {
	err := storage.LoadTasks("data/tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	http.HandleFunc("/task", handlers.TaskHandler)
	http.HandleFunc("/task/", handlers.TaskByIDHandler)

	fmt.Println("Start server")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
