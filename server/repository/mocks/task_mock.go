package repository

import (
	"server/models"
	"server/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTaskRepository houses logic to retrieve tasks from a mock repository
type MockTaskRepository struct{}

// NewMockTaskRepository convience function to create a MockTaskRepository
func NewMockTaskRepository() repository.TaskRepository {
	return &MockTaskRepository{}
}

var taskId, _ = primitive.ObjectIDFromHex("60b9718cf735e63dae12b5d5")
var tasks = map[string]primitive.M{
	"1": {
		"ID":          taskId,
		"Description": "mocked_task",
		"Status":      true,
	},
}

// GetAll get all tasks from the repository
func (r MockTaskRepository) GetAll() primitive.M {
	var results []primitive.M
	for _, task := range tasks {
		results = append(results, task)
	}
	return primitive.M{
		"error":  false,
		"result": results,
	}
}

// Create a Task to the repository
func (r MockTaskRepository) Create(task models.Task) primitive.M {
	task.ID = primitive.NewObjectID()
	tasks[task.ID.Hex()] = primitive.M{
		"ID":          task.ID,
		"Description": task.Description,
		"Status":      task.Status,
	}
	return primitive.M{
		"error": false,
	}
}

// Complete a Task from the repository
func (r MockTaskRepository) Complete(taskId string) primitive.M {
	var newTask = tasks[taskId]
	newTask["Status"] = true
	delete(tasks, taskId)
	tasks[taskId] = newTask
	return primitive.M{
		"error": false,
	}
}

// Undo a Task from the repository
func (r MockTaskRepository) Undo(taskId string) primitive.M {
	var newTask = tasks[taskId]
	newTask["Status"] = false
	delete(tasks, taskId)
	tasks[taskId] = newTask
	return primitive.M{
		"error": false,
	}
}

// Delete a Task from the repository
func (r MockTaskRepository) Delete(taskId string) primitive.M {
	delete(tasks, taskId)
	return primitive.M{
		"error": false,
	}
}

// MockErroringTaskRepository returns errors for all operations.
type MockErroringTaskRepository struct{}

// NewMockErroringTaskRepository convience function to create a MockErroringTaskRepository
func NewMockErroringTaskRepository() repository.TaskRepository {
	return &MockErroringTaskRepository{}
}

// GetAll get all tasks from the repository
func (r MockErroringTaskRepository) GetAll() primitive.M {
	return primitive.M{
		"error": true,
	}
}

// Create a Task to the repository
func (r MockErroringTaskRepository) Create(task models.Task) primitive.M {
	return primitive.M{
		"error": true,
	}
}

// Delete a Task from the repository
func (r MockErroringTaskRepository) Complete(taskId string) primitive.M {
	return primitive.M{
		"error": true,
	}
}

// Delete a Task from the repository
func (r MockErroringTaskRepository) Undo(taskId string) primitive.M {
	return primitive.M{
		"error": true,
	}
}

// Delete a Task from the repository
func (r MockErroringTaskRepository) Delete(taskId string) primitive.M {
	return primitive.M{
		"error": true,
	}
}
