package entities

import "github.com/ortizdavid/go-rest-concepts/config"

func SetupMigrations() {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&User{})
}