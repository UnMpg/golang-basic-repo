package usecase

import (
	"go-project/api/repository"
	"go-project/models"

	// serviceemail "go-project/service/serviceEmail"
	"go-project/utils/encript"
	"go-project/utils/log"
	"time"
)

type UserUsecase struct {
	URepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return UserUsecase{URepository: userRepo}
}

func (Uuc *UserUsecase) RegisterUser(data string) error {
	return nil
}

func (Uc *UserUsecase) CreateUser(req models.User) (models.DataUserCreate, error) {
	var data models.DataUserCreate

	hashPassword, err := encript.HashPassword(req.Password)
	if err != nil {
		log.Log.Error(log.Register, "Error Hashed Password", err)
		return data, err
	}
	now := time.Now()
	newUser := models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashPassword,
		Role:      "user",
		Verified:  "INACTIVE",
		CreatedAt: now,
	}

	detail := models.Detail{
		UserID:  "293924923",
		Address: "padang",
		Phone:   "082373873473483",
	}

	transaction := Uc.URepository.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if err := Uc.URepository.AddDetail(&detail); err != nil {
		transaction.Rollback()
		log.Log.Error(log.Register, "Error Save DB Detail", err)
		// return data, err
	}

	if err := Uc.URepository.RegisterUser(&newUser); err != nil {
		transaction.Rollback()
		log.Log.Error(log.Register, "Error Save DB User", err)
		// return data, err
	}

	// if err := serviceemail.SendEmailRegister(newUser, models.EmailData{URL: "", FirstName: "fendy", Subject: "Registrasion"}); err != nil {
	// 	return data, err
	// }

	if err := transaction.Commit().Error; err != nil {
		log.Log.Error(log.Register, "Error committing transaction", err)
		return data, err
	}
	data = models.DataUserCreate{
		Name:  newUser.Name,
		Email: newUser.Email,
		Role:  newUser.Role,
	}

	return data, nil
}

func (U *UserUsecase) CobaUsecaseTesting() (*models.User, error) {

	return U.URepository.CobaTestRepo()
}
