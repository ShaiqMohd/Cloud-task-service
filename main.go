package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = make(map[int]Task)
var currentID = 1

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	//CRUD Logic

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var task Task

			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			task.ID = currentID
			tasks[currentID] = task
			currentID++

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}

		if r.Method == http.MethodGet {
			idStr := r.URL.Query().Get("id")

			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "Invalid ID", http.StatusBadRequest)
					return
				}

				task, exists := tasks[id]
				if !exists {
					http.Error(w, "Task not found", http.StatusNotFound)
					return
				}

				w.Header().Set("Content-Type", "appplication/json")
				json.NewEncoder(w).Encode(task)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks)
			return
		}

		if r.Method == http.MethodDelete {
			idstr := r.URL.Query().Get("id")

			if idstr == "" {
				http.Error(w, "Task ID required", http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(idstr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
			}

			_, exists := tasks[id]
			if !exists {
				http.Error(w, "Task not found", http.StatusNotFound)
			}

			delete(tasks, id)

			w.Write([]byte("Task deleted"))
			return
		}

		if r.Method == http.MethodPut {
			idstr := r.URL.Query().Get("id")

			if idstr == "" {
				http.Error(w, "Task ID required", http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(idstr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}

			task, exists := tasks[id]
			if !exists {
				http.Error(w, "Task not found", http.StatusNotFound)
				return
			}

			var updatedTask Task
			err = json.NewDecoder(r.Body).Decode(&updatedTask)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			task.Title = updatedTask.Title
			task.Done = updatedTask.Done

			tasks[id] = task

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	})

	http.ListenAndServe(":8080", nil)
}
