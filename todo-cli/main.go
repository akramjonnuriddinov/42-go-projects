package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id        int
	Title     string
	Completed bool
}

var tasks []Task
var ID int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n****Choose an option****")
		fmt.Println("1. Create task")
		fmt.Println("2. Read task")
		fmt.Println("3. Update task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		option := input(scanner, "\nEnter your option:")
		switch option {
		case "1":
			title := input(scanner, "Enter task name: ")
			createTask(title)
		case "2":
			readTask()
		case "3":
			id_ := getInutID(scanner)
			title_ := input(scanner, "Enter new name: ")
			completed_ := input(scanner, "Enter status (yes/no): ")
			updateTask(id_, title_, completed_)
		case "4":
			id_ := getInutID(scanner)
			deleteTask(id_)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Try again!")
		}
	}
}

func input(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	val := strings.TrimSpace(scanner.Text())
	if val == "" {
		fmt.Println("Input cannot be empty. Please try again")
		return input(scanner, prompt)
	}
	return val
}

func getNextId() int {
	ID++
	return ID - 1
}

func parseID(id_ string) (int, error) {
	id, err := strconv.Atoi(id_)
	if err != nil {
		fmt.Println("Invalid ID format.")
	}
	return id, nil
}

func getInutID(scanner *bufio.Scanner) string {
	return input(scanner, "Enter the ID: ")
}

func createTask(title string) {
	task := Task{
		Id:        getNextId(),
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, task)
	fmt.Println("Task added Successfully!")
}

func readTask() {
	if len(tasks) == 0 {
		fmt.Println("There's no task available!")
		return
	}
	fmt.Println("Tasks: ")
	for _, task := range tasks {
		fmt.Printf("%d. %s [%t] \n", task.Id, task.Title, task.Completed)
	}
}

func updateTask(id_ string, title_ string, completed_ string) {
	id, err := parseID(id_)
	if err != nil {
		return
	}
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Title = title_
			if completed_ == "yes" {
				tasks[i].Completed = true
			} else {
				tasks[i].Completed = false
			}
			fmt.Println("Task updated successfully!")
			return
		}
	}
	fmt.Println("No task found with the given ID.")
}

func deleteTask(id_ string) {
	id, err := parseID(id_)
	if err != nil {
		return
	}
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("No task found with the given ID.")
}
