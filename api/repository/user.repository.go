package repository

import (
	"go-project/models"
	"go-project/utils/log"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var Data string

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r *UserRepository) RegisterUser(newUser *models.User) error {
	result := r.DB.Create(&newUser)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate key Value uniq") {
		log.Log.Error(log.Register, "Error Create Data")
		return result.Error
	} else if result.Error != nil {
		log.Log.Error(log.Register, "Error Create Data")
		return result.Error
	}

	return nil
}
