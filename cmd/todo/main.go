package main

import (
	"TODO/internal/handlers"
	"TODO/internal/storage"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := os.Getenv("TODO_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	tasksFile := os.Getenv("TODO_TASKS_FILE")
	if tasksFile == "" {
		tasksFile = "data/tasks.json"
	}

	handlers.SetTasksFile(tasksFile)

	err := storage.LoadTasks(tasksFile)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	http.HandleFunc("/task", handlers.TaskHandler)
	http.HandleFunc("/task/", handlers.TaskByIDHandler)

	fmt.Println("Start server")

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      nil,
	}

	// Run the server in a goroutine so main can wait for shutdown signals.
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Error starting the server:", err)
		}
	}()

	// Wait until the process receives Ctrl+C or a termination signal.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down server...")

	// Give active requests a short window to finish before stopping the server.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		fmt.Println("Error shutting down server:", err)
	}
}
