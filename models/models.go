package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	UserID uint   `json:"-" gorm:"not null"`
	Notes  string `json:"notes" gorm:"text";not null;default:null`
}
