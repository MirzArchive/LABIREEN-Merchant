package repositories

import (
	"labireen-merchant/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(merchant *entities.Merchant) error
	GetById(id uuid.UUID) (*entities.Merchant, error)
	GetWhere(param string, email string) (*entities.Merchant, error)
	Update(merchant *entities.Merchant) error
	Delete(merchant *entities.Merchant) error
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db}
}

func (r *authRepositoryImpl) Create(merchant *entities.Merchant) error {
	return r.db.Create(&merchant).Error
}

func (r *authRepositoryImpl) GetById(id uuid.UUID) (*entities.Merchant, error) {
	var merchant entities.Merchant
	if err := r.db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (r *authRepositoryImpl) GetWhere(param string, args string) (*entities.Merchant, error) {
	var merchant entities.Merchant
	if err := r.db.Where(param+" = ?", args).First(&merchant).Error; err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (r *authRepositoryImpl) Update(merchant *entities.Merchant) error {
	return r.db.Save(merchant).Error
}

func (r *authRepositoryImpl) Delete(merchant *entities.Merchant) error {
	return r.db.Delete(merchant).Error
}
