package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/repository"

	"github.com/gorilla/mux"
)

type TaskController struct {
	taskRepository repository.TaskRepository
}

func NewTaskController(repo repository.TaskRepository) *TaskController {
	return &TaskController{repo}
}

func (t TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	payload := t.taskRepository.Create(task)
	json.NewEncoder(w).Encode(payload)
}

func (t TaskController) GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := t.taskRepository.GetAll()
	json.NewEncoder(w).Encode(payload)
}

func (t TaskController) CompleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	payload := t.taskRepository.Complete(params["id"])
	json.NewEncoder(w).Encode(payload)
}

func (t TaskController) UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	payload := t.taskRepository.Undo(params["id"])
	json.NewEncoder(w).Encode(payload)
}

func (t TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	payload := t.taskRepository.Delete(params["id"])
	json.NewEncoder(w).Encode(payload)
}
