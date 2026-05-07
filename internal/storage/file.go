package storage

import (
	"encoding/json"
	"os"
)

func LoadTasks(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		return err
	}
	return nil

}

func SaveTasks(path string) error {
	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
