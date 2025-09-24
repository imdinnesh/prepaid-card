package repository

import (
	models "github.com/imdinnesh/prepaid-card/services/auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByMobile(phone string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}


func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByMobile(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("mobile = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}