package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Name        string           `gorm:"type:varchar(255);unique_index"`
	Description string           `gorm:"type:text;unique_index"`
	Type        int              `gorm:"type:varchar(128);unique_index"`
	SignInStart time.Time        `gorm:"type:timestamp"`
	SignInEnd   time.Time        `gorm:"type:timestamp"`
	Form        *ApplicationForm
}
