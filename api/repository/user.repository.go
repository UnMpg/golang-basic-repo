package repository

import (
	"context"
	"go-project/models"
	"go-project/utils/log"
	"strings"

	"go.elastic.co/apm"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var Data string

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r *UserRepository) RegisterUser(ctx context.Context, newUser *models.User) error {
	span, _ := apm.StartSpan(ctx, "insertUser", "repository")
	defer span.End()
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

func (r *UserRepository) AddDetail(detail *models.Detail) error {
	add := r.DB.Create(&detail)
	if add.Error != nil {
		log.Log.Error(log.Register, "Error Create Data")
		return add.Error
	}

	return nil
}

func (r *UserRepository) CobaTestRepo() (*models.User, error) {
	data := models.User{
		Name:  "namsms",
		Email: "email@gmail.com",
	}
	return &data, nil
}
