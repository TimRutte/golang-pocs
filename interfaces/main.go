package main

import (
	"fmt"
	"log"
)

type Task interface {
	Execute(action string) error
}

type TaskProcessor struct {
	ID     int
	Worker Task
}

func (tp *TaskProcessor) PerformTask(action string) {
	if err := tp.Worker.Execute(action); err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Printf("Task completed!\nID: %d\nAction: %s\n", tp.ID, action)
	}
}

type HandlerA struct{}

func (ha HandlerA) Execute(action string) error {
	if action == "" {
		return fmt.Errorf("HandlerA: action cannot be empty")
	}
	log.Println("HandlerA processed the action.")
	return nil
}

type HandlerB struct{}

func (hb HandlerB) Execute(action string) error {
	if action == "" {
		return fmt.Errorf("HandlerB: action cannot be empty")
	}
	log.Println("HandlerB processed the action.")
	return nil
}

func main() {
	// Define the workers
	handlerA := HandlerA{}
	handlerB := HandlerB{}

	// Create processors
	processor1 := TaskProcessor{ID: 1, Worker: handlerA}
	processor2 := TaskProcessor{ID: 2, Worker: handlerB}

	// Execute tasks
	processor1.PerformTask("Process Task A")
	processor2.PerformTask("Process Task B")
}
