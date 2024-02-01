package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key" `
	Name             string    `json:"name" gorm:"type:varchar(255);not null" validate:"required,ReqStringNumberChar"`
	Email            string    `json:"email" gorm:"uniqueIndex;type:varchar(255);not null" validate:"required,ReqEmail"`
	Password         string    `json:"password" gorm:"type:varchar(255);not null" validate:"ReqStringNumberChar"`
	Role             string    `json:"role" gorm:"type:varchar(255);not null"`
	VerificationCode string    `json:"verificationCode" gorm:"type:varchar(255);"`
	Verified         string    `json:"verified" gorm:"type:varchar(255);not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Detail struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key" `
	UserID    string    `json:"userId" gorm:"type:varchar(255);"`
	Address   string    `json:"address" gorm:"type:varchar(255);"`
	NIK       string    `json:"nik" gorm:"type:varchar(255);" `
	Phone     string    `json:"phone" gorm:"type:varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataUserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
