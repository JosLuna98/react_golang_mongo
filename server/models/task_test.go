package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestIsEmpty(t *testing.T) {
	task := Task{}
	assert.True(t, task.IsEmpty())
}

func TestIsEmptyIdPresent(t *testing.T) {
	task := Task{
		ID: primitive.NewObjectID(),
	}
	assert.False(t, task.IsEmpty())
}

func TestIsEmptyDescriptionPresent(t *testing.T) {
	task := Task{
		Description: "testing_task",
	}
	assert.False(t, task.IsEmpty())
}

func TestIsEmptyAllFields(t *testing.T) {
	task := Task{
		ID:          primitive.NewObjectID(),
		Description: "testing_task",
		Status:      true,
	}
	assert.False(t, task.IsEmpty())
}
