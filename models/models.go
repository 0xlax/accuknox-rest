package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	Notes string `json:"notes" gorm:"text";not null;default:null`
}
