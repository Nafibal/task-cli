package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func addTask(desc string) {
	tasks, err := readTasks()
	if err != nil {
		log.Fatal("error reading file")
	}

	task := Task{
		ID:          generateTaskID(tasks),
		Description: desc,
		Status:      statusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	tasks = append(tasks, task)

	err = writeTasks(tasks)
	if err != nil {
		log.Fatal("error writing file")
	}

	fmt.Printf("Task added successfully (ID: %v\n)", task.ID)
}

func deleteTask(id string) {
	tasks, err := readTasks()
	if err != nil {
		log.Fatal("error reading file")
	}
	var filteredTasks []Task

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("cannot convert string id to int")
	}

	for _, task := range tasks {
		if task.ID != intID {
			filteredTasks = append(filteredTasks, task)
		}
	}

	err = writeTasks(filteredTasks)
	if err != nil {
		log.Fatal("error writing file")
	}

	fmt.Printf("Task deleted successfully (ID: %v", intID)
}

func updateTask(id, desc, mark string) {
	tasks, err := readTasks()
	if err != nil {
		log.Fatal("error reading file")
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("cannot convert string id to int")
	}

	updatedTime := time.Now()
	idFound := false

	for i, task := range tasks {
		if task.ID == intID {
			idFound = true
			if desc != "" {
				tasks[i].Description = desc
			}
			if mark != "" {
				tasks[i].Status = mark
			}
			tasks[i].UpdatedAt = &updatedTime
		}
	}

	if !idFound {
		fmt.Println("no task with this id")
		return
	}

	err = writeTasks(tasks)
	if err != nil {
		log.Fatal("error writing file")
	}

	fmt.Printf("Task updated successfully (ID: %v", intID)
}

func listTask(status string) {
	tasks, err := readTasks()
	if err != nil {
		log.Fatal("error reading file")
	}

	if status != "" && status != statusTodo && status != statusDone && status != statusInProgress {
		fmt.Println("Invalid status filter. Use: todo | in-progress | done")
		return
	}

	var filteredTasks []Task
	if status != "" {
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else {
		filteredTasks = tasks
	}

	if len(filteredTasks) == 0 {
		fmt.Println("No Task Found")
		return
	}

	fmt.Println("---------------------------------------------------")
	fmt.Printf("%-5s %-30s %-12s\n", "ID", "Description", "Status")
	fmt.Println("---------------------------------------------------")

	// Print each task
	for i, task := range filteredTasks {
		fmt.Printf("%-5d %-30s %-12s\n", i, task.Description, task.Status)
	}
	fmt.Println("---------------------------------------------------")
}
