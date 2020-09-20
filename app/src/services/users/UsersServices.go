package users

import (
	"github.com/yoshi0202/grass-app/app/src/services"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Users []User

func FindUser(param int) *User {
	u := new(User)
	db := services.ConnectGorm()
	db.First(&u, param)
	return u
}

func FindAllUser() *Users {
	u := new(Users)
	db := services.ConnectGorm()
	db.Find(&u)
	return u
}
