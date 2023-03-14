package repositories

import (
	"labireen-merchant/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(Merchant *entities.Merchant) error
	GetById(id uuid.UUID) (*entities.Merchant, error)
	GetWhere(param string, email string) (*entities.Merchant, error)
	Update(Merchant *entities.Merchant) error
	Delete(Merchant *entities.Merchant) error
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db}
}

func (r *authRepositoryImpl) Create(Merchant *entities.Merchant) error {
	return r.db.Create(&Merchant).Error
}

func (r *authRepositoryImpl) GetById(id uuid.UUID) (*entities.Merchant, error) {
	var Merchant entities.Merchant
	if err := r.db.First(&Merchant, id).Error; err != nil {
		return nil, err
	}

	return &Merchant, nil
}

func (r *authRepositoryImpl) GetWhere(param string, args string) (*entities.Merchant, error) {
	var Merchant entities.Merchant
	if err := r.db.Where(param+" = ?", args).First(&Merchant).Error; err != nil {
		return nil, err
	}

	return &Merchant, nil
}

func (r *authRepositoryImpl) Update(Merchant *entities.Merchant) error {
	return r.db.Save(Merchant).Error
}

func (r *authRepositoryImpl) Delete(Merchant *entities.Merchant) error {
	return r.db.Delete(Merchant).Error
}
