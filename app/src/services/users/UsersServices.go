package users

import (
	"encoding/json"

	"github.com/yoshi0202/grass-app/app/src/services"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Users []User

func Find(param int) string {
	u := new(User)
	db := services.ConnectGorm()
	db.First(&u, param)
	json, _ := json.Marshal(u)
	return string(json)
}

func FindAll() string {
	u := new(Users)
	db := services.ConnectGorm()
	db.Find(&u)
	json, _ := json.Marshal(u)
	return string(json)
}
