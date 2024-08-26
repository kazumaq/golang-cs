package main

import (
	"fmt"
)

// Task represents a single task
type Task struct {
	ID          int
	Description string
	Done        bool
}

// TaskList represents a list of tasks
type TaskList struct {
	tasks []Task
	mu    sync.Mutex
}

// TaskManager interface defines methods for managing tasks
type TaskManager interface {
	AddTask(description string) int
	MarkAsDone(id int) error
	ListTasks()
}

// AddTask adds a new task to the task list
func (tl *TaskList) AddTask(description string) int {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	newTask := Task{
		ID:          len(tl.tasks) + 1,
		Description: description,
		Done:        false,
	}
	tl.tasks = append(tl.tasks, newTask)
	return newTask.ID
}

// MarkAsDone marks a task as done
func (tl *TaskList) MarkAsDone(id int) error {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	for i, task := range tl.tasks {
		if task.ID == id {
			tl.tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// ListTasks prints all tasks
func (tl *TaskList) ListTasks() {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	fmt.Println("Task List:")
	for _, task := range tl.tasks {
		status := "Not Done"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Description, status)
	}
}

// simulateWork simulates doing work on tasks concurrently
func simulateWork(tm TaskManager, taskDesc string, done chan<- bool) {
	id := tm.AddTask(taskDesc)
	fmt.Printf("Started task: %s\n", taskDesc)

	// Simulate work being done
	time.Sleep(time.Duration(2+id) * time.Second)

	err := tm.MarkAsDone(id)
	if err != nil {
		fmt.Printf("Error marking task as done: %v\n", err)
	} else {
		fmt.Printf("Finished task: %s\n", taskDesc)
	}
	done <- true
}

func main() {
	taskList := &TaskList{}

	// Demonstrate that TaskList implements TaskManager interface
	var _ TaskManager = taskList

	// Create some tasks concurrently
	taskDescriptions := []string{
		"Write sample Go program",
		"Demonstrate structs and methods",
		"Show interface usage",
		"Implement error handling",
		"Use goroutines for concurrency",
	}

	done := make(chan bool)
	for _, desc := range taskDescriptions {
		go simulateWork(taskList, desc, done)
	}

	// Wait for all tasks to complete
	for range taskDescriptions {
		<-done
	}

	// List all tasks
	taskList.ListTasks()
}
