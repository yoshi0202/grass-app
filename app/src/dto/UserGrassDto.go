package dto

import (
	"encoding/json"

	"gorm.io/gorm"
)

type UserGrass struct {
	gorm.Model
	LoginID string `json:"loginID"`
	GrassID uint   `json:"grassID"`
}

type UserGrasses []UserGrass

func NewUserGrass() *UserGrass {
	return &UserGrass{}
}

func NewUserGrasses() *UserGrasses {
	return &UserGrasses{}
}

func (g *UserGrass) ToJSON() string {
	json, _ := json.Marshal(g)
	return string(json)
}

func (g *UserGrasses) ToJSON() string {
	json, _ := json.Marshal(g)
	return string(json)
}
