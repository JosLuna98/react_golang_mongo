package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description,omitempty"`
	Status      bool               `json:"status,omitempty"`
}

type TaskList []Task

func (task Task) IsEmpty() bool {
	return task.ID.IsZero() && task.Description == ""
}

type TaskNotFoundError struct {
	Message string
}

func (task TaskNotFoundError) Error() string {
	return task.Message
}
