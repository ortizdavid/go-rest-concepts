package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-rest-concepts/entities"
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
	router.HandleFunc("POST /tasks/create-default", th.CreateDefaultTasks)
	router.Handle("GET /protected-apikey", ApiKeyAuthMiddleware(http.HandlerFunc(th.protectedApiKey)))
}


func (th TaskHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	tasks, err := taskModel.FindAll()
	count := len(tasks)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		WriteError(w, "Tasks not found", http.StatusNotFound)
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
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !taskModel.ExistsById(id) {
		WriteError(w, "Task with id: "+rawId+" not exists", http.StatusNotFound)
		return
	}
	WriteJSON(w, r, "Task found", http.StatusOK, task, 1)
}


func (th TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	var newTask entities.Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := taskModel.Create(newTask); err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJSON(w, r, "Task created", http.StatusCreated, newTask, 1)
}


func (th TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a Task")
}


func (th TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a Task")
}


func (th TaskHandler) CreateDefaultTasks(w http.ResponseWriter, r *http.Request) {
	count, _ := models.TaskModel{}.CreateDefault()
	WriteJSON(w, r, "Tasks added created", http.StatusCreated, nil, count)
}

func (th TaskHandler) protectedApiKey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Accessed API Key Resource")
}