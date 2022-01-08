package repository

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

type UserS struct {
	db *gorm.DB
}

func (rootRepo *Repositories) User() UserS {
	return UserS{db: rootRepo.Db}
}

func (u UserS) Login(username string) (models.User, error) {
	var user models.User
	err := u.db.Where("user_name = ?", username).Find(&user).Error
	return user, err
}

func (u UserS) DuplicateUserName(username string) int64 {
	var count int64
	u.db.Model(&models.User{}).Where("user_name = ?", username).Count(&count)
	return count
}

func (u UserS) Insert(user *models.User) error {
	err := u.db.Save(user).Error
	return err
}
