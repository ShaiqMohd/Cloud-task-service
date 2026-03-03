package main

import (
	"net/http"
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

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			task := Task{
				ID:    currentID,
				Title: "Sample Task",
				Done:  false,
			}
			tasks[currentID] = task
			currentID++
			w.Write([]byte("Task Created"))
			return
		}
		w.Write([]byte("Method Not Allowed"))
	})

	http.ListenAndServe(":8080", nil)
}
