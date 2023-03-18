package entities

import (
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	ID               uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	Name             string    `gorm:"not null"`
	Email            string    `gorm:"uniqueIndex;size:320;not null"`
	Password         string    `gorm:"not null"`
	Photo            string    `gorm:"not null"`
	VerificationCode string    `gorm:"not null"`
	Verified         bool      `gorm:"default:false"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

type MerchantRegister struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=8"`
	PasswordConfirm  string `json:"password_confirm" binding:"required"`
	VerificationCode string `json:"verification_code"`
}

type MerchantLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type MerchantRequest struct {
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
