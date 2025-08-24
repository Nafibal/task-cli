package main

import (
	"fmt"
	"os"
)

const (
	statusTodo       = "todo"
	statusInProgress = "in-progress"
	statusDone       = "done"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [args]")
		os.Exit(1)
	}

	command := os.Args[1]
	arg := os.Args[2:]

	switch command {
	case "add":
		if len(arg) != 1 {
			fmt.Println("Usage: task-cli add <task>")
			os.Exit(1)
		}
		addTask(arg[0])
	case "delete":
		if len(arg) != 1 {
			fmt.Println("Usage: task-cli delete <id>")
			os.Exit(1)
		}
		deleteTask(arg[0])
	case "update":
		if len(arg) < 2 {
			fmt.Println("Usage: task-cli update <id> <new-task>")
			os.Exit(1)
		}
		updateTask(arg[0], arg[1], "")
	case "mark-in-progress":
		if len(arg) != 1 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			os.Exit(1)
		}
		updateTask(arg[0], "", statusInProgress)
	case "mark-done":
		if len(arg) != 1 {
			fmt.Println("Usage: task-cli mark-done <id>")
			os.Exit(1)
		}
		updateTask(arg[0], "", statusDone)
	case "list":
		if len(arg) > 1 {
			fmt.Println("Usage: task-cli list [status]")
			os.Exit(1)
		}
		status := ""
		if len(arg) == 1 {
			status = arg[0]
		}
		listTask(status)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Usage: task-cli <add|delete|update|mark-in-progress|mark-done|list> [args]")
		os.Exit(1)
	}
}
