package users

import (
	"encoding/json"

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

func FindByGithubID(param string) *dto.User {
	u := dto.NewUser()
	db := services.ConnectGorm()
	db.Where(dto.User{Login: param}).Find(&u)
	return u
}

func FindOrCreate(user *dto.OAuthUser) *dto.User {
	us := dto.NewUser()
	db := services.ConnectGorm()
	us.Login = user.Login
	us.Name = user.Login
	us.Avatar = user.AvatarURL
	db.Where(dto.User{Login: user.Login}).FirstOrCreate(&us)
	return us
}

func GetUserIDByJSON(params []byte) *dto.OAuthUser {
	us := dto.NewOAuthUser()
	json.Unmarshal(params, &us)
	return us
}
