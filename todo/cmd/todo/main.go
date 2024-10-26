package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todo/internal/storage"
	"example.com/todo/internal/task"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter choice (1: Add, 2: List, 3: Update Status, 4: Delete, 5: Exit): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			taskDesc, _ := reader.ReadString('\n')
			taskDesc = strings.TrimSpace(taskDesc)
			fmt.Print("Is it done (yes or no): ")
			isDoneStr, _ := reader.ReadString('\n')
			isDoneStr = strings.TrimSpace(strings.ToLower(isDoneStr))
			isDone := (isDoneStr == "yes")

			todo := task.New(taskDesc, isDone) // Create a pointer to Todo
			storage.Add(todo)

		case "2":
			todos := storage.GetAll()
			if len(todos) == 0 {
				fmt.Println("No todos available.")
			} else {
				for i, todo := range todos {
					status := "not done"
					if todo.IsDone {
						status = "done"
					}
					fmt.Printf("%d: %s is %s\n", i, todo.Task, status) // Access fields of the struct
				}
			}

		case "3":
			todos := storage.GetAll()
			if len(todos) == 0 {
				fmt.Println("No todos available.")
				continue
			}
			fmt.Println("Current todos:")
			for i, todo := range todos {
				status := "not done"
				if todo.IsDone {
					status = "done"
				}
				fmt.Printf("%d: %s is %s\n", i, todo.Task, status)
			}

			fmt.Print("Enter the task number to update its status: ")
			taskNumStr, _ := reader.ReadString('\n')
			taskNumStr = strings.TrimSpace(taskNumStr)

			var taskNum int
			_, err := fmt.Sscanf(taskNumStr, "%d", &taskNum)
			if err != nil || taskNum < 0 || taskNum >= len(todos) {
				fmt.Println("Invalid task number.")
				continue
			}

			fmt.Print("Is it done (yes or no): ")
			isDoneStr, _ := reader.ReadString('\n')
			isDoneStr = strings.TrimSpace(strings.ToLower(isDoneStr))
			isDone := (isDoneStr == "yes")

			// Update the isDone status directly via the pointer
			todos[taskNum].IsDone = isDone

		case "4":
			todos := storage.GetAll()
			if len(todos) == 0 {
				fmt.Println("No todos available.")
				continue
			}
			fmt.Println("Current todos:")
			for i, todo := range todos {
				status := "not done"
				if todo.IsDone {
					status = "done"
				}
				fmt.Printf("%d: %s is %s\n", i, todo.Task, status)
			}

			fmt.Print("Enter the task number to delete it: ")
			taskNumStr, _ := reader.ReadString('\n')
			taskNumStr = strings.TrimSpace(taskNumStr)

			var taskNum int
			_, err := fmt.Sscanf(taskNumStr, "%d", &taskNum)
			if err != nil || taskNum < 0 || taskNum >= len(todos) {
				fmt.Println("Invalid task number.")
				continue
			} else {
				storage.DeleteElement(taskNum)
				fmt.Println("Task deleted successfully.")
			}

		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
		}
	}
}
