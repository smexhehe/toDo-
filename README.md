# Todo API

Simple REST API for managing todo tasks written in Go.

## Features

- Create, read, update and delete tasks
- JSON responses
- JSON error format
- Request validation
- In-memory storage
- Unit tests for handlers and storage

## Tech Stack

- Go
- net/http
- testify

## Run

```bash
go run ./cmd/todo
```

## Test

```bash
go test ./...
```

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | /task | Get all tasks |
| POST | /task | Create task |
| GET | /task/{id} | Get task by ID |
| PUT | /task/{id} | Update task |
| DELETE | /task/{id} | Delete task |

## Example Request

```json
{
  "title": "Learn Go",
  "done": false
}
```

## Error Response

```json
{
  "error": "title is required"
}
```

