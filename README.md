# To-Do List Manager

# Project Overview

The To-Do List Manager is a command-line application written in Go that allows users to manage tasks. Users can add new tasks, list all tasks, mark tasks as done, and delete tasks.

# Features:

- Add tasks with a description.

- List all tasks with their status (Pending or Done).

- Mark tasks as completed.

- Delete tasks by ID.

- Interactive command-line interface with help instructions.

# Usage Instructions

To use the To-Do List Manager, follow these steps:

1. Run the application (see the "How to Build and Run the Application" section below).

2. You will see a welcome message: Welcome to the To-Do List Manager! Type 'help' for instructions.

3. Use the following commands to interact with the program:

 - add: Add a new task by providing a description when prompted.

- list: List all tasks with their IDs and current status.

- done <task_id>: Mark a task as done using its ID.

- delete <task_id>: Delete a task using its ID.

- help: Display the list of available commands and their usage.

- exit: Exit the application.

4. Enter the desired command and follow the on-screen instructions to manage your tasks.

# How to Build and Run the Application

# Prerequisites

- Install latest version of Go. https://go.dev/dl/

Steps to Build and Run

1. Clone the repository:

git clone https://github.com/stihlshadow/DMQ_Module_1.git
cd DMQ_Module_1

2. Run the program:

go run main.go

3. To run unit tests on the program:

go test