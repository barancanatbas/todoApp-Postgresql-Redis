package repository

import (
	"todo/internal/config"

	"gorm.io/gorm"
)

type Repositories struct {
	Db *gorm.DB
}

var Repo *Repositories

func Set() {
	Repo = &Repositories{
		Db: config.Database,
	}
}

func Get() *Repositories {
	return Repo
}
