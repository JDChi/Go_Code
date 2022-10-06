package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar;not null"`
	Telephone string `gorm:"type:varchar;not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
