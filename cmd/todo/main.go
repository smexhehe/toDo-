package main

import (
	"TODO/internal/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/task", handlers.TaskHandler)
	http.HandleFunc("/task/", handlers.TaskByIDHandler)

	fmt.Println("Start server")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
