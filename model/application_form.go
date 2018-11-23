package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ApplicationForm struct {
	gorm.Model
	Name            string    `gorm:"type:varchar(255);unique_index"`
	Introduction    string    `gorm:"type:text;unique_index"`
	BackgroundImage string    `gorm:"type:text;unique_index"`
	Description     string    `gorm:"type:text;unique_index"`
	URL             string    `gorm:"type:text;unique_index"`
	Type            int       `gorm:"type:varchar(128);unique_index"`
	EffectiveStart  time.Time `gorm:"type:timestamp"`
	EffectiveEnd    time.Time `gorm:"type:timestamp"`
	UIs             []UI
}
