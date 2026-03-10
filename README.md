# Cloud Task Service (Go REST API)

A simple RESTful Task Management API built using **Go's net/http package**.
This project demonstrates backend fundamentals such as HTTP routing, JSON request/response handling, CRUD operations, and in-memory data storage.

The project is part of my learning journey in **Backend Development, Cloud, and DevOps**.

---

Features

* Health check endpoint
* Create tasks using JSON input
* Retrieve all tasks
* Retrieve a specific task by ID
* Delete tasks
* Proper HTTP status responses
* JSON request and response handling


Tech Stack

Language: Go
HTTP Server: net/http
Data Format: JSON
Storage: In-memory (map)


## API Endpoints

### Health Check

```
GET /health
```

Response:

```
OK
```

---

### Create Task

```
POST /tasks
```

Request Body:

```json
{
  "title": "Study Go",
  "done": false
}
```

Response:

```json
{
  "id": 1,
  "title": "Study Go",
  "done": false
}
```

---

### Get All Tasks

```
GET /tasks
```

Response:

```json
{
  "1": {
    "id": 1,
    "title": "Study Go",
    "done": false
  }
}
```

---

### Get Task by ID

```
GET /tasks?id=1
```

Response:

```json
{
  "id": 1,
  "title": "Study Go",
  "done": false
}
```

---

### Delete Task

```
DELETE /tasks?id=1
```

Response:

```
Task deleted
```

---

## Running the Project

Clone the repository:

```
git clone https://github.com/YOUR_USERNAME/cloud-task-service.git
```

Navigate to the project folder:

```
cd cloud-task-service
```

Run the server:

```
go run main.go
```

The server will start on:

```
http://localhost:8080
```

---

## Example Workflow

1. Start the server
2. Create tasks using Postman
3. Retrieve tasks using browser or Postman
4. Delete tasks via DELETE endpoint

---

## Future Improvements

Planned improvements for this project:

* Add **task update endpoint (PUT/PATCH)**
* Add **persistent storage with PostgreSQL**
* Containerize the application using **Docker**
* Deploy the API to **AWS**
* Add **CI/CD pipeline with GitHub Actions**
* Provision infrastructure using **Terraform**

---

## Learning Goals

This project is part of my preparation for roles in:

* Backend Development
* Cloud Engineering
* DevOps Engineering

It focuses on building strong fundamentals in API design and cloud-native development.

---
