# Cloud Task Service

A RESTful Task Management API built with Go, PostgreSQL, and Docker.

This project demonstrates backend development fundamentals including REST APIs, CRUD operations, database integration, environment-based configuration, containerization, and service orchestration using Docker Compose.

## Features

* Create tasks
* Retrieve all tasks
* Retrieve a task by ID
* Update tasks
* Delete tasks
* PostgreSQL persistence
* Dockerized application
* Docker Compose multi-container setup
* Environment variable configuration
* Health check endpoint

## Tech Stack

* Go
* PostgreSQL
* Docker
* Docker Compose
* REST API
* JSON

## Project Structure

```text
cloud-task-service
│
├── handlers
├── models
├── storage
│   └── db.go
├── Dockerfile
├── docker-compose.yml
├── .gitignore
├── main.go
└── README.md
```

## API Endpoints

### Health Check

```http
GET /health
```

### Create Task

```http
POST /tasks
```

Request:

```json
{
  "title": "Study Go",
  "done": false
}
```

### Get All Tasks

```http
GET /tasks
```

### Get Task By ID

```http
GET /tasks?id=1
```

### Update Task

```http
PUT /tasks?id=1
```

Request:

```json
{
  "title": "Study Docker",
  "done": true
}
```

### Delete Task

```http
DELETE /tasks?id=1
```

## Running the Project

### Using Docker Compose

```bash
docker compose up --build
```

Application:

```text
http://localhost:8080
```

### Environment Variables

Create a `.env` file:

```env
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=tasksdb
DB_HOST=db
DB_PORT=5432
```

## Architecture

```text
Client
  ↓
Go REST API
  ↓
Storage Layer
  ↓
PostgreSQL
```

## Learning Objectives

This project was built to strengthen skills in:

* Backend Development
* Database Integration
* Docker & Containerization
* Cloud Engineering
* DevOps Fundamentals

## Future Improvements

* AWS Deployment
* GitHub Actions CI/CD
* Terraform Infrastructure
* Authentication & Authorization
* API Documentation
* Automated Testing

## Deployment & CI/CD

- Dockerized using Docker and Docker Compose
- Deployed on AWS EC2
- Configured Jenkins Pipeline for automated build and deployment
- Managed environment variables using .env
