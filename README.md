# Todo API

![CI](https://github.com/smexhehe/toDo-/actions/workflows/ci.yml/badge.svg)

Simple REST API for managing todo tasks written in Go.

## Features

- Create, read, update and delete tasks
- JSON responses
- JSON error format
- Request validation
- File-backed JSON storage
- Tasks are loaded on startup and saved after changes
- Unit tests for handlers and storage
- HTTP server timeouts
- Graceful shutdown on Ctrl+C
- Configuration via environment variables



## Tech Stack

- Go
- net/http
- testify

## Run

```bash
go run ./cmd/todo
```
## Configuration

The server can be configured with environment variables:

| Variable | Default | Description |
|---|---|---|
| TODO_ADDR | :8080 | HTTP server address |
| TODO_TASKS_FILE | data/tasks.json | Path to JSON storage file |

PowerShell example:

```powershell
$env:TODO_ADDR=":3000"
$env:TODO_TASKS_FILE="data/dev-tasks.json"
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

## Docker

Build image:

```bash
docker build -t todo-api .
```

Run container:

```bash
docker run --rm -p 8080:8080 todo-api
```

Run with local data persistence:

```bash
docker run --rm -p 8080:8080 -v ${PWD}/data:/app/data todo-api
```


## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | /task | Get all tasks |
| POST | /task | Create task |
| GET | /task/{id} | Get task by ID |
| PUT | /task/{id} | Replace task |
| PATCH | /task/{id} | Partially update task |
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

