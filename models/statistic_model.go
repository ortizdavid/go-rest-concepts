package models

type StatisticCount struct {
	Users 				int64
	Tasks 				int64
}

func GetStatisticsCount() StatisticCount {
	countUsers, _ :=  UserModel{}.Count()
	countTasks, _ := TaskModel{}.Count()

	return StatisticCount{
		Users:           countUsers,
		Tasks:           countTasks,
	}
}
