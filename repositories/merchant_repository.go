package repositories

import (
	"labireen-merchant/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	Create(merchant *entities.Merchant) error
	GetById(id uuid.UUID) (*entities.Merchant, error)
	GetWhere(param string, email string) (*entities.Merchant, error)
	Update(merchant *entities.Merchant) error
	Delete(merchant *entities.Merchant) error
}

type merchantRepositoryImpl struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepositoryImpl{db}
}

func (rp *merchantRepositoryImpl) Create(merchant *entities.Merchant) error {
	return rp.db.Create(&merchant).Error
}

func (rp *merchantRepositoryImpl) GetById(id uuid.UUID) (*entities.Merchant, error) {
	var merchant entities.Merchant
	if err := rp.db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (rp *merchantRepositoryImpl) GetWhere(param string, args string) (*entities.Merchant, error) {
	var merchant entities.Merchant
	if err := rp.db.Where(param+" = ?", args).First(&merchant).Error; err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (rp *merchantRepositoryImpl) Update(merchant *entities.Merchant) error {
	return rp.db.Save(merchant).Error
}

func (rp *merchantRepositoryImpl) Delete(merchant *entities.Merchant) error {
	return rp.db.Delete(merchant).Error
}
