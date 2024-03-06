package handlers

import (
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-rest-concepts/models"
)

type TaskHandler struct {
}

func (th TaskHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /tasks", th.getAllTasks)
	router.HandleFunc("GET /tasks/{id}", th.getTask)
	router.HandleFunc("POST /tasks", th.createTask)
	router.HandleFunc("PUT /tasks/{id}", th.updateTask)
	router.HandleFunc("DELETE /tasks/{id}", th.deleteTask)
}


func (th TaskHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.TaskModel{}.FindAll()
	count := len(tasks)
	if err != nil {
		WriteError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		WriteError(w, r, "Tasks not found", http.StatusNotFound)
		return
	}
	WriteJSON(w, r, "Tasks Found Sucessfuly", http.StatusOK, tasks, count)

}


func (th TaskHandler) getTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	rawId := r.PathValue("id")
	id := conversion.StringToInt(rawId)
	task, err := taskModel.FindById(id)
	if err != nil {
		WriteError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	if !taskModel.ExistsById(id) {
		WriteError(w, r, "Task with id: "+rawId+" not exists", http.StatusNotFound)
		return
	}
	WriteJSON(w, r, "Task found", http.StatusOK, task, 1)
}


func (th TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, r, "Task created", http.StatusCreated, nil, 1)
}


func (th TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a Task")
}


func (th TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a Task")
}