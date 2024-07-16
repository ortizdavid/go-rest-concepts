package models

import "gorm.io/gorm"

type StatisticCount struct {
	Users 				int64
	Tasks 				int64
}

func NewStatisticsCount(db *gorm.DB) StatisticCount {
	countUsers, _ :=  NewUserModel(db).Count()
	countTasks, _ := NewTaskModel(db).Count()

	return StatisticCount{
		Users:           countUsers,
		Tasks:           countTasks,
	}
}
