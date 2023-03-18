package services

import (
	"errors"
	"labireen-merchant/entities"
	"labireen-merchant/pkg/crypto"
	"labireen-merchant/repositories"

	"github.com/google/uuid"
)

type AuthService interface {
	RegisterCustomer(merchant entities.CustomerRegister) error
	LoginCustomer(merchant entities.CustomerLogin) (uuid.UUID, error)
	VerifyCustomer(email string) error
}

type authServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authServiceImpl{repo}
}

func (asr *authServiceImpl) RegisterCustomer(merchant entities.CustomerRegister) error {
	hashedPassword, err := crypto.HashValue(merchant.Password)
	if err != nil {
		return errors.New("failed to encrypt given data")
	}

	assignID, err := uuid.NewRandom()
	if err != nil {
		return errors.New("failed to assign unique uuid")
	}

	user := entities.Merchant{
		ID:               assignID,
		Name:             merchant.Name,
		Email:            merchant.Email,
		Password:         hashedPassword,
		VerificationCode: merchant.VerificationCode,
	}

	err = asr.repo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

func (asr *authServiceImpl) LoginCustomer(merchant entities.CustomerLogin) (uuid.UUID, error) {
	user, err := asr.repo.GetWhere("email", merchant.Email)
	if err != nil {
		return uuid.UUID{}, errors.New("user not found")
	}

	if !user.Verified {
		return uuid.UUID{}, errors.New("user has not verified")
	}

	if err := crypto.CheckHash(merchant.Password, user.Password); err != nil {
		return uuid.UUID{}, errors.New("password is not valid or incorrect")
	}

	return user.ID, nil
}

func (asr *authServiceImpl) VerifyCustomer(code string) error {
	user, err := asr.repo.GetWhere("verification_code", code)
	if err != nil {
		return errors.New("user not found")
	}

	user.VerificationCode = ""
	user.Verified = true

	if err := asr.repo.Update(user); err != nil {
		return errors.New("failed to update user data")
	}

	return nil
}
