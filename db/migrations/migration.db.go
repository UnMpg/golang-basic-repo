package main

import (
	"fmt"
	"go-project/config"
	"go-project/models"
	"go-project/utils/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	var err error
	configEnv := config.AppEnv
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", configEnv.DBHost, configEnv.DBUserName, configEnv.DBUserPassword, configEnv.DBName, configEnv.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Log.Error("Failed connection to database")
	}

}

func init() {
	ConnectionDB()
}
func main() {
	if err := DB.AutoMigrate(&models.User{}, &models.Detail{}); err != nil {
		log.Log.Error("error to migrate table", err.Error())
	}
}
