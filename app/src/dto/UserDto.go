package dto

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Users []User

func NewUser() *User {
	return &User{}
}

func NewUsers() *Users {
	return &Users{}
}

func (u *User) ToJSON() string {
	json, _ := json.Marshal(u)
	return string(json)
}

func (u *Users) ToJSON() datatypes.JSON {
	json, _ := json.Marshal(u)
	return json
}
