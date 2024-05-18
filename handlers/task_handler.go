package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/serialization"
	"github.com/ortizdavid/go-rest-concepts/entities"
	"github.com/ortizdavid/go-rest-concepts/models"
)


type TaskHandler struct {
}


func (th TaskHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /api/tasks", th.getAllTasks)
	router.HandleFunc("GET /api/tasks/{id}", th.getTask)
	router.HandleFunc("POST /api/tasks", th.createTask)
	router.HandleFunc("PUT /api/tasks/{id}", th.updateTask)
	router.HandleFunc("DELETE /api/tasks/{id}", th.deleteTask)
	router.HandleFunc("POST /api/tasks/import-csv", th.importTasksCSV)

	router.HandleFunc("GET /api/tasks-xml", th.getAllTasksXml)
	router.HandleFunc("GET /api/tasks-xml/{id}", th.getTaskXml)
}


func (th TaskHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	tasks, err := taskModel.FindAll()
	count := len(tasks)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		httputils.WriteJsonError(w, "Tasks not found", http.StatusNotFound)
		return
	}
	httputils.WriteJson(w, http.StatusOK, tasks, count)
}


func (th TaskHandler) getTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	rawId := r.PathValue("id")
	id := conversion.StringToInt(rawId)
	task, err := taskModel.FindById(id)

	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusOK, task)
}


func (th TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	var task entities.Task

	if err := serialization.DecodeJson(r.Body, &task); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Uuid, created, updated
	task.UniqueId = encryption.GenerateUUID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_, err := taskModel.Create(task)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusCreated, task)
}


func (th TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	var updatedTask entities.Task

	id := r.PathValue("id")
	taskId := conversion.StringToInt(id)
	task, err := taskModel.FindById(taskId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := serialization.DecodeJson(r.Body, &updatedTask); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return 
	}

	task.TaskName = updatedTask.TaskName
	task.UserId = updatedTask.UserId
	task.StartDate = task.EndDate
	task.EndDate = updatedTask.EndDate
	task.Description = updatedTask.Description
	task.UpdatedAt = time.Now()
	
	_, err = taskModel.Update(task)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusOK, task)
}


func (th TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	id := r.PathValue("id")
	taskId := conversion.StringToInt(id)
	task, err := taskModel.FindById(taskId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = taskModel.Delete(task)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusNoContent, nil)
	fmt.Fprintf(w, "Delete a task")
}


func (th TaskHandler) getAllTasksXml(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	tasks, err := taskModel.FindAll()
	count := len(tasks)
	if err != nil {
		httputils.WriteXmlError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		httputils.WriteXmlError(w, "Tasks not found", http.StatusNotFound)
		return
	}
	httputils.WriteXml(w, http.StatusOK, tasks, count)
}


func (th TaskHandler) getTaskXml(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	rawId := r.PathValue("id")
	id := conversion.StringToInt(rawId)

	task, err := taskModel.FindById(id)
	if err != nil {
		httputils.WriteXmlError(w, err.Error(), http.StatusNotFound)
		return
	}
	
	exists, _ := taskModel.ExistsRecord("task_id", id)
	if !exists {
		httputils.WriteXmlError(w, "Task with id: "+rawId+" not exists", http.StatusNotFound)
		return
	}
	httputils.WriteXmlSimple(w, http.StatusOK, task)
}



func (th TaskHandler) importTasksCSV(w http.ResponseWriter, r *http.Request) {
	var taskModel models.TaskModel
	
	csvFile, _, err := r.FormFile("csv_file")
	if err != nil {
		httputils.WriteJsonError(w, "could not upload csv", http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	if err := models.SkipCsvHeader(reader); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasksCsv, err := models.ParseTaskFromCSV(reader)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = taskModel.CreateBatch(tasksCsv)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	httputils.WriteJson(w, http.StatusCreated, nil, len(tasksCsv))
}

