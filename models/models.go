package models

import "gorm.io/gorm"

type word struct {
	gorm.Models
	Question string `json:"question" gorm:"text";not null;default:null`
	Answer   string `json:"answer" gorm:"text";not null;default:null`
}
