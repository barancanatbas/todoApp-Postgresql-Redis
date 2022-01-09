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

func (t TodoS) GetTodo(todoId uint, userId uint) (models.Todo, error) {
	var todo models.Todo

	err := t.db.
		Where("id = ?", todoId).
		Where("userfk = ?", userId).
		Where("completed = false").
		Find(&todo).Error
	return todo, err
}

func (t TodoS) Update(todo *models.Todo) error {
	err := t.db.Save(todo).Error
	return err
}

func (t TodoS) Delete(todoId uint, userId uint) error {
	err := t.db.Model(&models.Todo{}).
		Where("id = ?", todoId).
		Where("userfk = ?", userId).
		Delete(&models.Todo{}).Error

	return err
}
