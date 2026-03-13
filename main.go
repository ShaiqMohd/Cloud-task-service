package main

import (
	"net/http"

	"cloud-task-service/handlers"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	//CRUD Logic

	http.HandleFunc("/tasks", handlers.TaskHandler)

	http.ListenAndServe(":8080", nil)
}
