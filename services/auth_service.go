package services

import (
	"errors"
	"labireen-merchant/entities"
	"labireen-merchant/pkg/crypto"
	"labireen-merchant/pkg/mail"
	"labireen-merchant/repositories"
	"os"

	"github.com/google/uuid"
)

type AuthService interface {
	Registermerchant(merchant *entities.MerchantRegister) (mail.EmailData, error)
	Loginmerchant(merchant entities.MerchantLogin) (uuid.UUID, error)
	Verifymerchant(email string) error
	FindByParams(param string, args string) (entities.Merchant, error)
	Updatemerchant(merchant entities.MerchantRequest) error
	ResetPassword(pwd entities.MerchantReset, id uuid.UUID) error
}

type authServiceImpl struct {
	rp repositories.MerchantRepository
}

func NewAuthService(rp repositories.MerchantRepository) AuthService {
	return &authServiceImpl{rp}
}

func (svc *authServiceImpl) Registermerchant(merchant *entities.MerchantRegister) (mail.EmailData, error) {
	if merchant.Password != merchant.PasswordConfirm {
		return mail.EmailData{}, errors.New("password mismatch")
	}

	hashedPassword, err := crypto.HashValue(merchant.Password)
	if err != nil {
		return mail.EmailData{}, errors.New("failed to encrypt given data")
	}

	assignID, err := uuid.NewRandom()
	if err != nil {
		return mail.EmailData{}, errors.New("failed to assign unique uuid")
	}

	user := entities.Merchant{
		ID:               assignID,
		Name:             merchant.Name,
		Email:            merchant.Email,
		Password:         hashedPassword,
		VerificationCode: crypto.Encode(merchant.Email),
	}

	err = svc.rp.Create(&user)
	if err != nil {
		return mail.EmailData{}, err
	}

	email := mail.EmailData{
		Email:   []string{user.Email},
		URL:     os.Getenv("EMAIL_CLIENT_ORIGIN") + "/auth/verify/" + user.VerificationCode,
		Name:    user.Name,
		Subject: "Your account verification code",
	}

	return email, nil
}

func (svc *authServiceImpl) Loginmerchant(merchant entities.MerchantLogin) (uuid.UUID, error) {
	user, err := svc.rp.GetWhere("email", merchant.Email)
	if err != nil {
		return uuid.Nil, errors.New("user not found")
	}

	if !user.Verified {
		return uuid.Nil, errors.New("user has not verified")
	}

	if err := crypto.CheckHash(merchant.Password, user.Password); err != nil {
		return uuid.Nil, errors.New("password is not valid or incorrect")
	}

	return user.ID, nil
}

func (svc *authServiceImpl) Verifymerchant(code string) error {
	user, err := svc.rp.GetWhere("verification_code", code)
	if err != nil {
		return errors.New("user not found")
	}

	user.VerificationCode = ""
	user.Verified = true

	if err := svc.rp.Update(user); err != nil {
		return errors.New("failed to update user data")
	}

	return nil
}

func (svc *authServiceImpl) FindByParams(param string, args string) (entities.Merchant, error) {
	user, err := svc.rp.GetWhere(param, args)
	if err != nil {
		return entities.Merchant{}, err
	}

	return *user, nil
}

func (svc *authServiceImpl) Updatemerchant(merchant entities.MerchantRequest) error {
	user, err := svc.FindByParams("email", merchant.Email)
	if err != nil {
		return err
	}

	user = entities.Merchant{
		Name:  merchant.Name,
		Email: merchant.Email,
		Photo: merchant.Photo,
	}

	if err := svc.rp.Update(&user); err != nil {
		return err
	}

	return nil
}

func (svc *authServiceImpl) ResetPassword(pwd entities.MerchantReset, id uuid.UUID) error {
	if pwd.Password != pwd.PasswordConfirm {
		return errors.New("password mismatch")
	}

	user, err := svc.rp.GetById(id)
	if err != nil {
		return err
	}

	hashedPassword, err := crypto.HashValue(pwd.Password)
	if err != nil {
		return errors.New("failed to encrypt given data")
	}

	user.Password = hashedPassword

	if err := svc.rp.Update(user); err != nil {
		return err
	}

	return nil
}
