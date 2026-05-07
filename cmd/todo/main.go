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
	err := storage.LoadTasks("data/tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	http.HandleFunc("/task", handlers.TaskHandler)
	http.HandleFunc("/task/", handlers.TaskByIDHandler)

	fmt.Println("Start server")
	// Создаём HTTP-сервер
	server := &http.Server{
		// Адрес и порт, которые будет слушать сервер
		Addr: ":8080",
		// Максимальное время чтения запроса от клиента
		ReadTimeout: 5 * time.Second,
		// Максимальное время записи ответа клиенту
		WriteTimeout: 10 * time.Second,
		// nil = использовать http.DefaultServeMux
		// (маршруты, зарегистрированные через http.HandleFunc)
		Handler: nil,
	}

	// Запускаем сервер в отдельной goroutine,
	// чтобы main goroutine могла продолжить выполнение
	go func() {
		// Запускаем HTTP-сервер
		// ListenAndServe блокирует выполнение,
		// пока сервер работает
		err := server.ListenAndServe()

		// Проверяем ошибку
		// http.ErrServerClosed возникает при нормальном Shutdown()
		// и не считается настоящей ошибкой
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Error starting the server:", err)
		}
	}()

	// Создаём канал для получения сигналов ОС
	stop := make(chan os.Signal, 1)

	// Подписываемся на сигналы:
	// os.Interrupt -> Ctrl+C
	// syscall.SIGTERM -> сигнал завершения процесса
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Блокируем main goroutine,
	// ждём сигнал завершения
	<-stop

	fmt.Println("Shutting down server...")

	// Создаём контекст с timeout 5 секунд
	// Shutdown будет ждать завершения активных запросов
	// максимум 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Освобождение ресурсов контекста
	defer cancel()

	// Graceful shutdown:
	//
	// 1. Перестаёт принимать новые подключения
	// 2. Ждёт завершения текущих запросов
	// 3. Закрывает соединения
	err = server.Shutdown(ctx)
	if err != nil {
		fmt.Println("Error shutting down server:", err)
	}

}
