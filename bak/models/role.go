package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name             string `gorm:"type:varchar(128);unique_index"`
	PermissionGroups []PermissionGroup
}
