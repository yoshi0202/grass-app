package usergrassess

import (
	"encoding/json"

	"github.com/yoshi0202/grass-app/app/src/services"
	"gorm.io/gorm"
)

type UserGrass struct {
	gorm.Model
	UserId  uint   `json:"userId"`
	GrassId string `json:"grassId"`
}

type UserGrassess []UserGrass

func FindByUserId(param int) string {
	ug := new(UserGrassess)
	db := services.ConnectGorm()
	db.Where("user_id = ?", param).Find(&ug)
	json, _ := json.Marshal(ug)
	return string(json)
}

func FindByGrassId(param int) string {
	ug := new(UserGrassess)
	db := services.ConnectGorm()
	db.Where("grass_id = ?", param).Find(&ug)
	json, _ := json.Marshal(ug)
	return string(json)
}
