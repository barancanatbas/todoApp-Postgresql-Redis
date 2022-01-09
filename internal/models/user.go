package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Surname  string
	Password string `json:"-"`
	UserName string
	Gmail    string
}
