package middleware

import (
	"server/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TgetAllTask() primitive.M {
	return getAllTask()
}

func TinsertOneTask(task models.ToDoList) primitive.M {
	return insertOneTask(task)
}

func TcompleteTask(taskId string) primitive.M {
	return completeTask(taskId)
}

func TundoTask(taskId string) primitive.M {
	return undoTask(taskId)
}

func TdeleteOneTask(taskId string) primitive.M {
	return deleteOneTask(taskId)
}