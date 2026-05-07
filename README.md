# Todo API

Simple REST API for managing todo tasks written in Go.

## Features

- Create, read, update and delete tasks
- JSON responses
- JSON error format
- Request validation
- File-backed JSON storage
- Tasks are loaded on startup and saved after changes
- Unit tests for handlers and storage

## Tech Stack

- Go
- net/http
- testify

## Run

```bash
go run ./cmd/todo
```

## Data Storage

Tasks are stored in:

```text
data/tasks.json
```

The file is ignored by Git because it contains local runtime data.

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

