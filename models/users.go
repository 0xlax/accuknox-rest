package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"text";not null;default:null`
	Email     string `json:"email" gorm:"text";not null;default:null`
	Password  string `json:"password" gorm:"text";not null;default:null`
	SessionID string `json:"sessionid" gorm:"text";not null;default:null`
	Notes     []Word `json:"notes" gorm:"foreignkey:UserID"`
}

type LoginUser struct {
	gorm.Model
	Email    string `json:"email" gorm:"text";not null;default:null`
	Password string `json:"password" gorm:"text";not null;default:null`
}
