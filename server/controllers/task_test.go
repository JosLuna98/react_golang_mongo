package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"server/models"
	mocks "server/repository/mocks"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllTasks(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	uc := NewTaskController(mocks.NewMockTaskRepository())
	uc.GetAllTask(w, r)
	resp := w.Result()

	bs, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	assert.Equal(t, "{\"error\":false,\"result\":[{\"Description\":\"mocked_task\",\"ID\":\"60b9718cf735e63dae12b5d5\",\"Status\":true}]}\n", string(bs))
}

func CreateTask(t *testing.T) {
	var taskId, _ = primitive.ObjectIDFromHex("60b9718cf735e63dae12b5d5")
	task := models.Task{
		ID:          taskId,
		Description: "mocked_task",
		Status:      true,
	}

	bs, _ := json.Marshal(&task)
	r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(bs))
	w := httptest.NewRecorder()

	uc := NewTaskController(mocks.NewMockTaskRepository())
	uc.CreateTask(w, r)
	resp := w.Result()

	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
}

func TestDeleteTask(t *testing.T) {
	r := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()
	uc := NewTaskController(mocks.NewMockTaskRepository())
	uc.DeleteTask(w, r)
	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
}
