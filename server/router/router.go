package router

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func Router(tc *controllers.TaskController) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", tc.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", tc.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", tc.CompleteTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", tc.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", tc.DeleteTask).Methods("DELETE", "OPTIONS")
	return router
}
