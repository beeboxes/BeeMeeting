package models

import "github.com/jinzhu/gorm"

type Permission struct {
	gorm.Model
	Name              string `gorm:"type:varchar(128);unique_index"`
	PermissionGroupID uint
}

type PermissionGroup struct {
	gorm.Model
	Name        string `gorm:"type:varchar(128);unique_index"`
	Permissions []Permission
}
