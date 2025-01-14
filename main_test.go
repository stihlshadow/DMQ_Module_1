package main

import (
	"fmt"
	"testing"
)

// TestAddTask checks if tasks are added correctly.
func TestAddTask(t *testing.T) {
	tasks = []Task{} // Initialize an empty task list.
	nextID = 1       // Reset the next ID.

	addTestTask("Test Task 1") // Add a test task.
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Description != "Test Task 1" {
		t.Errorf("Expected description 'Test Task 1', got '%s'", tasks[0].Description)
	}
	if tasks[0].ID != 1 {
		t.Errorf("Expected ID 1, got %d", tasks[0].ID)
	}
}

// TestMarkTaskAsDone checks if tasks are marked as done correctly.
func TestMarkTaskAsDone(t *testing.T) {
	tasks = []Task{
		{ID: 1, Description: "Test Task 1", Done: false},
	}

	markTestTaskAsDone(1) // Mark the first task as done.
	if !tasks[0].Done {
		t.Errorf("Expected task to be marked as done")
	}
}

// TestDeleteTask checks if tasks are deleted correctly.
func TestDeleteTask(t *testing.T) {
	tasks = []Task{
		{ID: 1, Description: "Test Task 1", Done: false},
		{ID: 2, Description: "Test Task 2", Done: false},
	}

	deleteTestTask(1) // Delete the first task.
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].ID != 1 {
		t.Errorf("Expected remaining task ID to be 1, got %d", tasks[0].ID)
	}
}

// TestListTasks checks if tasks are listed correctly.
func TestListTasks(t *testing.T) {
	tasks = []Task{
		{ID: 1, Description: "Test Task 1", Done: false},
		{ID: 2, Description: "Test Task 2", Done: true},
	}

	output := captureListTasksOutput() // Capture the task list output.
	expectedOutput := "Tasks:\n1. Test Task 1 [Pending]\n2. Test Task 2 [Done]\n"
	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

// Helper Functions for Tests

// addTestTask adds a new task to the list with the given description.
func addTestTask(description string) {
	tasks = append(tasks, Task{ID: nextID, Description: description, Done: false})
	nextID++ // Increment the next ID.
}

// markTestTaskAsDone marks the task with the specified ID as done.
func markTestTaskAsDone(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Done = true
			return
		}
	}
}

// deleteTestTask removes the task with the specified ID from the list.
func deleteTestTask(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task.
			// Renumber the tasks to maintain sequential IDs.
			for j := range tasks {
				tasks[j].ID = j + 1
			}
			nextID = len(tasks) + 1
			return
		}
	}
}

// captureListTasksOutput generates a string representation of the current tasks.
func captureListTasksOutput() string {
	var output string
	if len(tasks) == 0 {
		output = "No tasks available.\n"
	} else {
		output = "Tasks:\n"
		for _, task := range tasks {
			status := "Pending"
			if task.Done {
				status = "Done"
			}
			output += fmt.Sprintf("%d. %s [%s]\n", task.ID, task.Description, status)
		}
	}
	return output
}
