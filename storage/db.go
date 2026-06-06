package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"cloud-task-service/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Println("No .env file found")
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		err = DB.Ping()
		if err == nil {
			break
		}

		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostgreSQL")

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT,
		done BOOLEAN
	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTask(title string, done bool) (models.Task, error) {

	query := "INSERT INTO tasks (title, done) VALUES ($1, $2) RETURNING id"

	var id int

	err := DB.QueryRow(query, title, done).Scan(&id)
	if err != nil {
		return models.Task{}, err
	}

	return models.Task{
		ID:    id,
		Title: title,
		Done:  done,
	}, nil
}

func GetAllTasks() ([]models.Task, error) {

	rows, err := DB.Query("SELECT id, title, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task

		rows.Scan(&t.ID, &t.Title, &t.Done)

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func GetTaskByID(id int) (models.Task, error) {

	var task models.Task

	query := "SELECT id, title, done FROM tasks WHERE id=$1"

	err := DB.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Done,
	)

	return task, err
}

func DeleteTask(id int) error {

	query := "DELETE FROM tasks WHERE id=$1"

	_, err := DB.Exec(query, id)

	return err
}

func UpdateTask(id int, title string, done bool) error {

	query := "UPDATE tasks SET title=$1, done=$2 WHERE id=$3"

	_, err := DB.Exec(query, title, done, id)

	return err
}
