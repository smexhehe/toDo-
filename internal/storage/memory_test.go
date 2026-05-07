package storage

import (
	"TODO/internal/models"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTask(t *testing.T) {
	Tasks = nil

	task := models.Task{
		Title: "test",
		Done:  false,
	}

	createdTask := CreateTask(task)

	require.Equal(t, 1, createdTask.ID)
	require.Equal(t, "test", createdTask.Title)
	require.False(t, createdTask.Done)
	require.Len(t, Tasks, 1)
}

func TestGetTaskByID(t *testing.T) {
	Tasks = []models.Task{
		{ID: 1, Title: "first", Done: false},
		{ID: 2, Title: "second", Done: true},
	}

	task, ok := GetTaskByID(2)

	require.True(t, ok)
	require.Equal(t, 2, task.ID)
	require.Equal(t, "second", task.Title)
	require.True(t, task.Done)
}

func TestGetTaskByIDNotFound(t *testing.T) {
	Tasks = []models.Task{
		{ID: 1, Title: "first", Done: false},
	}

	task, ok := GetTaskByID(999)

	require.False(t, ok)
	require.Equal(t, models.Task{}, task)
}

func TestUpdateTask(t *testing.T) {
	Tasks = []models.Task{
		{ID: 1, Title: "old", Done: false},
	}

	updatedTask, ok := UpdateTask(1, models.Task{
		Title: "new",
		Done:  true,
	})

	require.True(t, ok)
	require.Equal(t, 1, updatedTask.ID)
	require.Equal(t, "new", updatedTask.Title)
	require.True(t, updatedTask.Done)

	require.Equal(t, "new", Tasks[0].Title)
	require.True(t, Tasks[0].Done)
}

func TestUpdateTaskNotFound(t *testing.T) {
	Tasks = []models.Task{
		{ID: 1, Title: "old", Done: false},
	}

	updatedTask, ok := UpdateTask(999, models.Task{
		Title: "new",
		Done:  true,
	})

	require.False(t, ok)
	require.Equal(t, models.Task{}, updatedTask)
	require.Equal(t, "old", Tasks[0].Title)
	require.False(t, Tasks[0].Done)
}

func TestDeleteTask(t *testing.T) {
	Tasks = []models.Task{
		{ID: 1, Title: "first", Done: false},
		{ID: 2, Title: "second", Done: true},
	}

	ok := DeleteTask(1)

	require.True(t, ok)
	require.Len(t, Tasks, 1)
	require.Equal(t, 2, Tasks[0].ID)
	require.Equal(t, "second", Tasks[0].Title)
}

func TestDeleteTaskNotFound(t *testing.T) {
	Tasks = []models.Task{
		{ID: 1, Title: "first", Done: false},
	}

	ok := DeleteTask(999)

	require.False(t, ok)
	require.Len(t, Tasks, 1)
	require.Equal(t, 1, Tasks[0].ID)
}

func TestLoadTasks(t *testing.T) {
	Tasks = nil

	file := t.TempDir() + "/tasks.json"

	err := os.WriteFile(
		file,
		[]byte(`[{"id":1,"title":"from file","done":true}]`),
		0644,
	)
	require.NoError(t, err)

	err = LoadTasks(file)
	require.NoError(t, err)

	require.Len(t, Tasks, 1)
	require.Equal(t, 1, Tasks[0].ID)
	require.Equal(t, "from file", Tasks[0].Title)
	require.True(t, Tasks[0].Done)
}

func TestSaveTasks(t *testing.T) {
	Tasks = []models.Task{
		{ID: 1, Title: "saved", Done: true},
	}

	file := t.TempDir() + "/tasks.json"

	err := SaveTasks(file)
	require.NoError(t, err)

	Tasks = nil

	err = LoadTasks(file)
	require.NoError(t, err)

	require.Len(t, Tasks, 1)
	require.Equal(t, 1, Tasks[0].ID)
	require.Equal(t, "saved", Tasks[0].Title)
	require.True(t, Tasks[0].Done)
}

func TestCreateTaskAfterDeleteUsesNextID(t *testing.T) {
	Tasks = []models.Task{
		{ID: 2, Title: "existing", Done: false},
	}

	createdTask := CreateTask(models.Task{
		Title: "new",
		Done:  false,
	})

	require.Equal(t, 3, createdTask.ID)
}
