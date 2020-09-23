package usergrassess

import (
	"encoding/json"
	"fmt"

	"github.com/yoshi0202/grass-app/app/src/dto"
	"github.com/yoshi0202/grass-app/app/src/services"
)

func Create(userID string, gr *dto.Grass) {
	ug := dto.NewUserGrass()
	ug.LoginID = userID
	ug.GrassID = gr.ID
	db := services.ConnectGorm()
	db.Create(&ug)
}
func FindByUserId(param int) string {
	ug := dto.NewUserGrass()
	db := services.ConnectGorm()
	db.Where("user_id = ?", param).Find(&ug)
	json, _ := json.Marshal(ug)
	return string(json)
}

func FindByGrassId(param int) string {
	ug := dto.NewUserGrass()
	db := services.ConnectGorm()
	db.Where("grass_id = ?", param).Find(&ug)
	json, _ := json.Marshal(ug)
	return string(json)
}

func FindAllByUserId(param string) *dto.UserGrasses {
	ug := dto.NewUserGrasses()
	db := services.ConnectGorm()
	db.Where("login_id = ?", param).Find(&ug)
	for _, v := range *ug {
		fmt.Println(v.LoginID)
		fmt.Println(v.GrassID)
	}

	return ug
}
