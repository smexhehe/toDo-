package storage

import "TODO/internal/models"

var Tasks []models.Task

func CreateTask(task models.Task) models.Task {
	task.ID = nextID()
	Tasks = append(Tasks, task)
	return task
}

func GetTaskByID(id int) (models.Task, bool) {
	for i := range Tasks {
		if Tasks[i].ID == id {
			return Tasks[i], true
		}
	}

	return models.Task{}, false
}

func UpdateTask(id int, task models.Task) (models.Task, bool) {
	for i := range Tasks {
		if Tasks[i].ID == id {
			task.ID = id
			Tasks[i] = task
			return task, true
		}
	}

	return models.Task{}, false
}

func DeleteTask(id int) bool {
	for i := range Tasks {
		if Tasks[i].ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return true
		}
	}

	return false
}

func GetTasks() []models.Task {
	return Tasks
}

func nextID() int {
	maxID := 0

	for i := range Tasks {
		if Tasks[i].ID > maxID {
			maxID = Tasks[i].ID
		}
	}

	return maxID + 1
}

func PatchTask(id int, patch models.TaskPatch) (models.Task, bool) {
	for i := range Tasks {
		if Tasks[i].ID == id {
			if patch.Title != nil {
				Tasks[i].Title = *patch.Title
			}
			if patch.Done != nil {
				Tasks[i].Done = *patch.Done
			}
			return Tasks[i], true
		}
	}
	return models.Task{}, false
}
