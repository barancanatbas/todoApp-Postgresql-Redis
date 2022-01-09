package models

import "gorm.io/gorm"

type ForgetPassword struct {
	gorm.Model
	Userfk uint   `gorm:"column:userfk" json:"userfk"`
	User   User   `gorm:"foreignkey:userfk"`
	Code   string `gorm:"type:CHAR(4)"`
}
