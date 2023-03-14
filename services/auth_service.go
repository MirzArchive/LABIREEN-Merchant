package services

import (
	"errors"
	"labireen-merchant/entities"
	"labireen-merchant/repositories"
	"labireen-merchant/utilities/crypto"

	"github.com/google/uuid"
)

type AuthService interface {
	RegisterMerchant(Merchant entities.MerchantRegister) error
	LoginMerchant(Merchant entities.MerchantLogin) (uuid.UUID, error)
	VerifyMerchant(email string) error
}

type authServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authServiceImpl{repo}
}

func (asr *authServiceImpl) RegisterMerchant(Merchant entities.MerchantRegister) error {
	hashedPassword, err := crypto.HashValue(Merchant.Password)
	if err != nil {
		return errors.New("failed to encrypt given data")
	}

	assignID, err := uuid.NewRandom()
	if err != nil {
		return errors.New("failed to assign unique uuid")
	}

	user := entities.Merchant{
		ID:               assignID,
		Name:             Merchant.Name,
		Email:            Merchant.Email,
		Password:         hashedPassword,
		PhoneNumber:      Merchant.PhoneNumber,
		VerificationCode: Merchant.VerificationCode,
	}

	err = asr.repo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

func (asr *authServiceImpl) LoginMerchant(Merchant entities.MerchantLogin) (uuid.UUID, error) {
	user, err := asr.repo.GetWhere("email", Merchant.Email)
	if err != nil {
		return uuid.UUID{}, errors.New("user not found")
	}

	if !user.Verified {
		return uuid.UUID{}, errors.New("user has not verified")
	}

	if err := crypto.CheckHash(Merchant.Password, user.Password); err != nil {
		return uuid.UUID{}, errors.New("password is not valid or incorrect")
	}

	return user.ID, nil
}

func (asr *authServiceImpl) VerifyMerchant(code string) error {
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
