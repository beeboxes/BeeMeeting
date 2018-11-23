package models

import "github.com/jinzhu/gorm"

type UI struct {
	gorm.Model
	Name string `gorm:"type:text;unique_index"`
	Type int
	Code string
}
