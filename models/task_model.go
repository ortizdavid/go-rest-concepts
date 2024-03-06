package models

import (
	"fmt"
	"time"
	"github.com/ortizdavid/go-rest-concepts/entities"
)

type TaskModel struct {
}

func (tm TaskModel) FindAll() ([]entities.Task, error) {
	return []entities.Task{
		{
			TaskId:    1,
			UserId:    101,
			TaskName:  "Task 1",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 7),
		},
		{
			TaskId:    2,
			UserId:    102,
			TaskName:  "Task 2",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 7),
		},
		{
			TaskId:    3,
			UserId:    103,
			TaskName:  "Task 3",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 7),
		},
		{
			TaskId:    4,
			UserId:    104,
			TaskName:  "Task 4",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 7),
		},
		{
			TaskId:    5,
			UserId:    105,
			TaskName:  "Task 5",
			StartDate: time.Now(),
			EndDate:   time.Now().AddDate(0, 0, 7),
		},
	}, nil
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