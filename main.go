package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var Tasks []Task

func taskHandler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(Tasks)
	if err != nil {
		fmt.Println("Error marshaling:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {

	http.HandleFunc("/task", taskHandler)

	fmt.Println("Start server")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
