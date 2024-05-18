package models

import (
	"time"

	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-rest-concepts/config"
	"github.com/ortizdavid/go-rest-concepts/entities"
	"gorm.io/gorm"
)

type UserModel struct {
	LastInsertId int
}

func (userModel *UserModel) Create(user entities.User) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	user.UniqueId = encryption.GenerateUUID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userModel.LastInsertId = user.UserId
	return result, nil
}

func (UserModel) FindAll() ([]entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) Update(user entities.User) (*gorm.DB, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	result := db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (UserModel) FindById(id int) (entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, id)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (UserModel) FindByUniqueId(uniqueId string) (entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (UserModel) FindByUserName(userName string) (entities.User, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.First(&user, "user_name=?", userName)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}


func (UserModel) Search(param interface{}) ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE user_name=?", param).Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) Count() (int64, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var count int64
	result := db.Table("users").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}


func (UserModel) GetDataById(id int) (entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var userData entities.UserData
	result := db.Raw("SELECT * FROM view_user_data WHERE user_id=?", id).Scan(&userData)
	if result.Error != nil {
		return entities.UserData{}, result.Error
	}
	return userData, nil
}


func (UserModel) FindAllData() ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	result := db.Raw("SELECT * FROM view_user_data").Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (UserModel) FindAllDataLimit(start int, end int) ([]entities.UserData, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var users []entities.UserData
	result := db.Raw("SELECT * FROM view_user_data LIMIT ?, ?", start, end).Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}


func (UserModel) ExistsRecord(fieldName string, value any) (bool, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var user entities.User
	result := db.Where(fieldName+" = ?", value).First(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return user.UserId != 0, nil
}
