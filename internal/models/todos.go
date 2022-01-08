package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Userfk uint `gorm:"column:userfk" json:"userfk"`
	User   User `gorm:"foreignkey:userfk"`
	Task   string
}
