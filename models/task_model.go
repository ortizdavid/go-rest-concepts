package models

import (
	"fmt"
	"github.com/ortizdavid/go-rest-concepts/entities"
)

type TaskModel struct {
}

var sliceTasks []entities.Task


func (tm TaskModel) FindAll() ([]entities.Task, error) {
	return sliceTasks, nil
}


func (tm TaskModel) Create(task entities.Task) error {
	sliceTasks = append(sliceTasks, task)
	return nil
}

func (tm TaskModel) CreateBatch(tasks []entities.Task) (int, error) {
	sliceTasks = append(sliceTasks, tasks...)
	return len(tasks), nil
}

func (tm TaskModel) CreateDefault() (int, error) {
	tasks := []entities.Task{
		{
			TaskId:    1,
			UserId:    101,
			TaskName:  "Task 1",
			StartDate: "2024-02-02 11:00:00",
			EndDate:   "2024-02-02 11:00:00",
		},
		{
			TaskId:    2,
			UserId:    102,
			TaskName:  "Task 2",
			StartDate: "2024-02-02 11:00:00",
			EndDate:   "2024-02-02 11:00:00",
		},
		{
			TaskId:    3,
			UserId:    103,
			TaskName:  "Task 3",
			StartDate: "2024-02-02 11:00:00",
			EndDate:   "2024-02-02 11:00:00",
		},
		{
			TaskId:    4,
			UserId:    104,
			TaskName:  "Task 4",
			StartDate: "2024-02-02 11:00:00",
			EndDate:   "2024-02-02 11:00:00",
		},
		{
			TaskId:    5,
			UserId:    105,
			TaskName:  "Task 5",
			StartDate: "2024-02-02 11:00:00",
			EndDate:   "2024-02-02 11:00:00",
		},
	}
	sliceTasks = append(sliceTasks, tasks...)
	return len(tasks), nil
}


func (tm TaskModel) FindById(id int) (entities.Task, error) {
	tasks, _ := tm.FindAll()
	for _, t := range tasks {
		if t.TaskId == id {
			return t, nil
		}
	}
	return entities.Task{}, fmt.Errorf("error fetching id: %d", id)
}


func (tm TaskModel) ExistsById(id int) bool {
	tasks, _ := tm.FindAll()
	for _, t := range tasks {
		if t.TaskId == id {
			return true
		}
	}
	return false
}