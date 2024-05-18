package models

import (
	"github.com/ortizdavid/go-rest-concepts/config"
	"github.com/ortizdavid/go-rest-concepts/entities"
	"gorm.io/gorm"
)

type TaskModel struct {
	LastInsertId 	int
}

func (taskModel *TaskModel) Create(task entities.Task) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	
	taskModel.LastInsertId = task.TaskId
	return result, nil
}

func (taskModel *TaskModel) CreateBatch(tasks []entities.Task) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	tx := db.Begin()
	result := db.Create(&tasks)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return result, nil
}

func (TaskModel) FindAll() ([]entities.Task, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.Task
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (TaskModel) Update(task entities.Task) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (TaskModel) Delete(task entities.Task) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Delete(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (TaskModel) FindById(id int) (entities.Task, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var task entities.Task
	result := db.First(&task, id)
	if result.Error != nil {
		return entities.Task{}, result.Error
	}
	return task, nil
}

func (TaskModel) FindByUniqueId(uniqueId string) (entities.Task, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var task entities.Task
	result := db.First(&task, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.Task{}, result.Error
	}
	return task, nil
}

func (TaskModel) Search(param string) ([]entities.TaskData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	result := db.Raw("SELECT * FROM view_task_data WHERE task_name=?", param).Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (TaskModel) GetDataById(id int) (entities.TaskData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var taskData entities.TaskData
	result := db.Raw("SELECT * FROM view_task_data WHERE task_id=?", id).Scan(&taskData)
	if result.Error != nil {
		return entities.TaskData{}, result.Error
	}
	return taskData, nil
}

func (TaskModel) GetDataByUniqueId(uniqueId string) (entities.TaskData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var taskData entities.TaskData
	result := db.Raw("SELECT * FROM view_task_data WHERE unique_id=?", uniqueId).Scan(&taskData)
	if result.Error != nil {
		return entities.TaskData{}, result.Error
	}
	return taskData, nil
}

func (TaskModel) FindAllData() ([]entities.TaskData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	result := db.Raw("SELECT * FROM view_task_data").Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (TaskModel) FindAllDataByUserIdLimit(userId int, start int, end int) ([]entities.TaskData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tasks []entities.TaskData
	result := db.Raw("SELECT * FROM view_task_data WHERE user_id=? LIMIT ?, ?", userId, start, end).Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (TaskModel) Count() (int64, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	result := db.Table("tasks").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (TaskModel) ExistsRecord(fieldName string, value any) (bool, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var task entities.Task
	result := db.Where(fieldName+" = ?", value).First(&task)
	if result.Error != nil {
		return false, result.Error
	}
	return task.TaskId != 0, nil
}
