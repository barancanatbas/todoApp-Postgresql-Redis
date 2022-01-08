package repository

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

type TodoS struct {
	db *gorm.DB
}

func (rootRepo *Repositories) Todo() TodoS {
	return TodoS{db: rootRepo.Db}
}

func (t TodoS) List(userId uint) ([]models.Todo, error) {
	var todos []models.Todo

	err := t.db.Where("userfk = ?", userId).Find(&todos).Error
	return todos, err
}

func (t TodoS) Insert(todo *models.Todo) error {
	err := t.db.Debug().Model(&models.Todo{}).Create(todo).Error
	return err
}
