package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string
	Email         string
	Password      string
	Token         string
	Refresh_Token string
	Photo         Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
