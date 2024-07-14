package repository

import (
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user entities.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetByUsername(username string) (entities.User, error) {
	var user entities.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
