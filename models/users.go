package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"text";not null;default:null`
	Email    string `json:"email" gorm:"text";not null;default:null`
	Password string `json:"password" gorm:"text";not null;default:null`
	Notes    []Word `json:"notes" gorm:"-"`
}

type LoginUser struct {
	gorm.Model
	Email    string `json:"email" gorm:"text";not null;default:null`
	Password string `json:"password" gorm:"text";not null;default:null`
}
