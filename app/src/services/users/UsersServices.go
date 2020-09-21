package users

import (
	"github.com/yoshi0202/grass-app/app/src/dto"
	"github.com/yoshi0202/grass-app/app/src/services"
)

func Find(param int) *dto.User {
	u := dto.NewUser()
	db := services.ConnectGorm()
	db.First(&u, param)
	return u
}

func FindAll() *dto.Users {
	us := dto.NewUsers()
	db := services.ConnectGorm()
	db.Find(&us)
	return us
}
