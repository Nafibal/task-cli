package main

import (
	"encoding/json"
	"io"
	"os"
)

func readTasks() ([]Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err == io.EOF {
		tasks = []Task{}
	} else if err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeTasks(tasks []Task) error {
	file, err := os.OpenFile("tasks.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

func generateTaskID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	maxID := tasks[0].ID

	for i := range tasks {
		if tasks[i].ID > maxID {
			maxID = tasks[i].ID
		}
	}

	return maxID + 1
}
