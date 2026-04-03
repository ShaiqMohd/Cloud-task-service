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

		query := "INSERT INTO tasks (title, done) VALUES ($1, $2) RETURNING id"

		var id int
		err = storage.DB.QueryRow(query, task.Title, task.Done).Scan(&id)
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}

		task.ID = id

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

			var task models.Task

			query := "SELECT id, title, done FROM tasks WHERE id=$1"
			err = storage.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Done)

			if err != nil {
				http.Error(w, "Task not found", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
		// Get all tasks
		rows, err := storage.DB.Query("SELECT id, title, done FROM tasks")
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tasks []models.Task

		for rows.Next() {
			var t models.Task
			rows.Scan(&t.ID, &t.Title, &t.Done)
			tasks = append(tasks, t)
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

		query := "DELETE FROM tasks WHERE id=$1"
		_, err = storage.DB.Exec(query, id)
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

		query := "UPDATE tasks SET title=$1, done=$2 WHERE id=$3"
		_, err = storage.DB.Exec(query, updatedTask.Title, updatedTask.Done, id)
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
