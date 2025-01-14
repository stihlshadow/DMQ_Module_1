package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a single to-do item with an ID, description, and completion status.
type Task struct {
	ID          int
	Description string
	Done        bool
}

var tasks []Task      // Slice to store all tasks.
var nextID int = 1    // Variable to keep track of the next task ID.

func main() {
	reader := bufio.NewReader(os.Stdin) // Reader for capturing user input.
	fmt.Println("Welcome to the To-Do List Manager! Type 'help' for instructions.")

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n') // Read user input.
		input = strings.TrimSpace(input)    // Remove extra spaces and newlines.
		
		switch {
		case input == "add":
			addTask(reader)
		case input == "list":
			listTasks()
		case strings.HasPrefix(input, "done"):
			markTaskAsDone(input)
		case strings.HasPrefix(input, "delete"):
			deleteTask(input)
		case input == "help":
			showHelp()
		case input == "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

// addTask prompts the user for a task description and adds a new task to the list.
func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n') // Read task description.
	description = strings.TrimSpace(description)
	if description == "" {
		fmt.Println("Task description cannot be empty.")
		return
	}
	tasks = append(tasks, Task{ID: nextID, Description: description, Done: false}) // Add the new task.
	nextID++ // Increment the next ID.
	fmt.Println("Task added successfully.")
}

// listTasks displays all tasks with their IDs and statuses.
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, status)
	}
}

// markTaskAsDone marks a task as completed based on the provided task ID.
func markTaskAsDone(input string) {
	parts := strings.Split(input, " ") // Split input into command and ID.
	if len(parts) != 2 {
		fmt.Println("Usage: done <task_id>")
		return
	}

	taskID, err := strconv.Atoi(parts[1]) // Convert ID to integer.
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			if task.Done {
				fmt.Println("Task is already marked as done.")
				return
			}
			tasks[i].Done = true // Mark the task as done.
			fmt.Println("Task marked as done.")
			return
		}
	}

	fmt.Println("Task not found.")
}

// deleteTask removes a task from the list based on the provided task ID.
func deleteTask(input string) {
	parts := strings.Split(input, " ") // Split input into command and ID.
	if len(parts) != 2 {
		fmt.Println("Usage: delete <task_id>")
		return
	}

	taskID, err := strconv.Atoi(parts[1]) // Convert ID to integer.
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task.
			fmt.Println("Task deleted successfully.")

			// Renumber the tasks to maintain sequential IDs.
			for j := range tasks {
				tasks[j].ID = j + 1
			}
			nextID = len(tasks) + 1
			return
		}
	}

	fmt.Println("Task not found.")
}

// showHelp displays the list of available commands and their usage.
func showHelp() {
	fmt.Print(`
Commands:
  add            - Add a new task
  list           - List all tasks
  done <task_id> - Mark a task as done
  delete <task_id> - Delete a task
  help           - Show this help menu
  exit           - Exit the application
`)
}
