package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cloud-task-service/models"
	"cloud-task-service/storage"
)

//CRUD Logic

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		task, err = storage.CreateTask(task.Title, task.Done)
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}

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

			task, err := storage.GetTaskByID(id)
			if err != nil {
				http.Error(w, "Task not found", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
		// Get all tasks
		tasks, err := storage.GetAllTasks()
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
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
			return
		}

		err = storage.DeleteTask(id)
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}

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

		var updatedTask models.Task
		err = json.NewDecoder(r.Body).Decode(&updatedTask)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err = storage.UpdateTask(
			id,
			updatedTask.Title,
			updatedTask.Done,
		)

		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Task updated"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed"))

}
