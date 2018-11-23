package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username           string `gorm:"type:varchar(128);unique_index"`
	FullName           string `gorm:"type:varchar(128);unique_index"`
	Email              string `gorm:"type:varchar(128);unique_index"`
	Phone              string `gorm:"type:varchar(32);unique_index"`
	EncryptedPassword  string `gorm:"type:varchar(512);unique_index"`
	ResetPasswordToken string `gorm:"type:varchar(512);unique_index"`
	Roles              []Role
}
