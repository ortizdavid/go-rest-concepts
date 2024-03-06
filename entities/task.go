package entities

import "time"

type Task struct {
	TaskId    	int
	UserId    	int
	TaskName  	string
	StartDate 	time.Time
	EndDate		time.Time
}