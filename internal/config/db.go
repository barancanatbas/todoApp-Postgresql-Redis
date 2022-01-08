package config

import (
	"todo/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	Database = connect()
	migrate()
}

func connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=mysql123 dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("veri tabanı bağlantısı başarısız.")
	}
	db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	return db

}

func migrate() {
	Database.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
}
