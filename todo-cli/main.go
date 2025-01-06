package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed string
}

var tasks []Task
var Id int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Create Task")
		fmt.Println("2. Read Task ")
		fmt.Println("3. Update Task")
		fmt.Println("4. Exit()")
		option := input(scanner, "\nEnter your option:")
		switch option {
		case "1":
			title := input(scanner, "Enter task name: ")
			createTask(title)
		case "2":
			readTask()
		case "3":
			ID := input(scanner, "Enter the ID: ")
			newTitle := input(scanner, "Enter new Name: ")
			newCompleted := input(scanner, "Enter the status(true/false): ")
			updateTask(newTitle, newCompleted, ID)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Try again: ")
		}
	}
}

func createTask(title string) {
	task := Task{
		ID:        Id,
		Title:     title,
		Completed: "false",
	}
	tasks = append(tasks, task)
	Id++
	fmt.Printf("\nTask created successfully! \n")
	fmt.Println()
}

func readTask() {
	if len(tasks) == 0 {
		fmt.Printf("\nThere's no task available \n")
	}
	fmt.Println("Tasks: ")
	for _, task := range tasks {
		fmt.Printf("%d.%s [%s] \n", task.ID, task.Title, task.Completed)
	}
	fmt.Println()
}

func updateTask(newTitle string, newCompleted string, ID string) {
	for i := 0; i < len(tasks); i++ {
		if ID == strconv.Itoa(tasks[i].ID) {
			tasks[i].Title = newTitle
			tasks[i].Completed = newCompleted
			fmt.Println("Task updated...")
			return
		}
	}
	fmt.Println("Task not found with given ID.")
}

func input(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
