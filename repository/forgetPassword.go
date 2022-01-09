package repository

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

type ForgetPass struct {
	db *gorm.DB
}

func (rootRepo *Repositories) ForgetPass() ForgetPass {
	return ForgetPass{db: rootRepo.Db}
}

func (f ForgetPass) Count(startDate string, endDate string, userId uint) (int, error) {
	var count int64

	err := f.db.Debug().Model(&models.ForgetPassword{}).Where("date(created_at) between ? and ?", startDate, endDate).Where("userfk = ?", userId).Count(&count).Error
	return int(count), err
}

func (f ForgetPass) Insert(forgetpass models.ForgetPassword) error {
	err := f.db.Create(&forgetpass).Error
	return err
}
