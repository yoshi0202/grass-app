package dto

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   uint
	Name string
}