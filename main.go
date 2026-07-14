package main

import (
	"fmt"
	"os"
	"time"
	"taskmanager/storage"
	"taskmanager/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("uso: todo <comando> [argumentos]")
		return
	}

	command := os.Args[1]

	tasks, err := storage.Load()
	if err != nil {
		fmt.Println("Erro ao carregar tarefas:", err)
		return
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("uso: todo add \"<título da tarefa>\"")
			return
		}
		newTask := task.Task{
			ID: nextID(tasks),
			Title: os.Args[2],
			Done: false,
			CreatedAt: time.Now(),
		}
		tasks = append(tasks, newTask)
		storage.Save(tasks)
		fmt.Println("Tarefa adicionada:", newTask.Title)

	
	case "list":
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d - %s\n", status, t.ID, t.Title)
		}

	default:
		fmt.Println("comando desconhecido", command)
	}
}

func nextID(tasks []task.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}