package services

import (
	"labireen-merchant/entities"
	"labireen-merchant/repositories"

	"github.com/google/uuid"
)

type MerchantService interface {
	UpdateMerchant(Merchant entities.MerchantRequest) error
	GetMerchant(id uuid.UUID) (entities.MerchantRequest, error)
}

type MerchantServiceImpl struct {
	repo repositories.AuthRepository
}

func NewMerchantService(repo repositories.AuthRepository) MerchantService {
	return &MerchantServiceImpl{repo}
}

func (csr *MerchantServiceImpl) UpdateMerchant(Merchant entities.MerchantRequest) error {
	return nil
}

func (csr *MerchantServiceImpl) GetMerchant(id uuid.UUID) (entities.MerchantRequest, error) {
	user, err := csr.repo.GetById(id)
	if err != nil {
		return entities.MerchantRequest{}, err
	}

	userResp := entities.MerchantRequest{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Photo:       user.Photo,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userResp, nil
}
