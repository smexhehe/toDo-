package handlers

import (
	"TODO/internal/models"
	"TODO/internal/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTasksFile(t *testing.T) {
	t.Helper()

	oldTasksFile := tasksFile
	tasksFile = t.TempDir() + "/tasks.json"

	t.Cleanup(func() {
		tasksFile = oldTasksFile
	})
}

func TestTaskHandlerCreateTask(t *testing.T) {
	setupTasksFile(t)
	storage.Tasks = nil

	req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(`{"title":"test","done":false}`))
	rr := httptest.NewRecorder()
	TaskHandler(rr, req)
	require.Equal(t, http.StatusCreated, rr.Code)

	var task models.Task
	err := json.NewDecoder(rr.Body).Decode(&task)
	require.NoError(t, err)

	require.Equal(t, 1, task.ID)
	require.Equal(t, "test", task.Title)
	require.False(t, task.Done)
	require.Len(t, storage.Tasks, 1)

}

func TestTaskHandlerCreateTaskEmptyTitle(t *testing.T) {
	storage.Tasks = nil

	req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(`{"title":"","done":false}`))
	rr := httptest.NewRecorder()
	TaskHandler(rr, req)
	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Len(t, storage.Tasks, 0)
	var response errorResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, "title is required", response.Error)
}

func TestTaskByIDHandlerGetTask(t *testing.T) {
	storage.Tasks = []models.Task{
		{ID: 1, Title: "test", Done: false},
	}
	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	rr := httptest.NewRecorder()
	TaskByIDHandler(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	var task models.Task
	err := json.NewDecoder(rr.Body).Decode(&task)
	require.NoError(t, err)
	require.Equal(t, 1, task.ID)
	require.Equal(t, "test", task.Title)
	require.False(t, task.Done)
}

func TestTaskByIDHandlerGetTaskNotFound(t *testing.T) {
	storage.Tasks = []models.Task{
		{ID: 1, Title: "test", Done: false},
	}
	req := httptest.NewRequest(http.MethodGet, "/task/999", nil)
	rr := httptest.NewRecorder()
	TaskByIDHandler(rr, req)
	require.Equal(t, http.StatusNotFound, rr.Code)
	var response errorResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, "id not found", response.Error)
}

func TestTaskByIDHandlerUpdateTask(t *testing.T) {
	setupTasksFile(t)
	storage.Tasks = []models.Task{
		{ID: 1, Title: "old", Done: false},
	}
	req := httptest.NewRequest(http.MethodPut, "/task/1", strings.NewReader(`{"title":"new","done":true}`))
	rr := httptest.NewRecorder()
	TaskByIDHandler(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	var task models.Task
	err := json.NewDecoder(rr.Body).Decode(&task)
	require.NoError(t, err)
	require.Equal(t, 1, task.ID)
	require.Equal(t, "new", task.Title)
	require.True(t, task.Done)
	require.Equal(t, "new", storage.Tasks[0].Title)
	require.True(t, storage.Tasks[0].Done)
}

func TestTaskByIDHandlerDeleteTask(t *testing.T) {
	setupTasksFile(t)
	storage.Tasks = []models.Task{
		{ID: 1, Title: "test", Done: false},
	}
	req := httptest.NewRequest(http.MethodDelete, "/task/1", nil)
	rr := httptest.NewRecorder()
	TaskByIDHandler(rr, req)
	require.Equal(t, http.StatusNoContent, rr.Code)
	require.Len(t, storage.Tasks, 0)
}

func TestTaskByIDHandlerPatchTaskDone(t *testing.T) {
	setupTasksFile(t)

	storage.Tasks = []models.Task{
		{ID: 1, Title: "old", Done: false},
	}

	req := httptest.NewRequest(http.MethodPatch, "/task/1", strings.NewReader(`{"done":true}`))
	rr := httptest.NewRecorder()

	TaskByIDHandler(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var task models.Task
	err := json.NewDecoder(rr.Body).Decode(&task)
	require.NoError(t, err)

	require.Equal(t, 1, task.ID)
	require.Equal(t, "old", task.Title)
	require.True(t, task.Done)

	require.Equal(t, "old", storage.Tasks[0].Title)
	require.True(t, storage.Tasks[0].Done)
}

func TestTaskByIDHandlerPatchTaskTitle(t *testing.T) {
	setupTasksFile(t)

	storage.Tasks = []models.Task{
		{ID: 1, Title: "old", Done: false},
	}

	req := httptest.NewRequest(http.MethodPatch, "/task/1", strings.NewReader(`{"title":"new"}`))
	rr := httptest.NewRecorder()

	TaskByIDHandler(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var task models.Task
	err := json.NewDecoder(rr.Body).Decode(&task)
	require.NoError(t, err)

	require.Equal(t, 1, task.ID)
	require.Equal(t, "new", task.Title)
	require.False(t, task.Done)

	require.Equal(t, "new", storage.Tasks[0].Title)
	require.False(t, storage.Tasks[0].Done)
}

func TestTaskByIDHandlerPatchTaskEmptyTitle(t *testing.T) {
	storage.Tasks = []models.Task{
		{ID: 1, Title: "old", Done: false},
	}

	req := httptest.NewRequest(http.MethodPatch, "/task/1", strings.NewReader(`{"title":""}`))
	rr := httptest.NewRecorder()

	TaskByIDHandler(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)

	var response errorResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, "title is required", response.Error)

	require.Equal(t, "old", storage.Tasks[0].Title)
	require.False(t, storage.Tasks[0].Done)
}

func TestTaskHandlerGetTasksEmpty(t *testing.T) {
	storage.Tasks = nil

	req := httptest.NewRequest(http.MethodGet, "/task", nil)
	rr := httptest.NewRecorder()

	TaskHandler(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.JSONEq(t, `[]`, rr.Body.String())
}
