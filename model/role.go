package model

type Role struct {
	Name             string `gorm:"type:varchar(128);unique_index"`
	PermissionGroups []PermissionGroup
}
