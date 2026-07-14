package storage

import (
	"encoding/json"
	"os"
	"taskmanager/task"
)
const fileName = "tasks.json"

func Load() ([]task.Task, error) {
	data, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		return []task.Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}