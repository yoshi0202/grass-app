package dto

import (
	"encoding/json"

	"github.com/yoshi0202/grass-app/app/src/constant"
)

type OAuth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Code         string `json:"code"`
}

func NewOAuth(code string) *OAuth {
	o := &OAuth{}
	c := constant.NewConst()
	o.ClientID = c.ClientID
	o.ClientSecret = c.ClientSecret
	o.RedirectURI = c.RedirectURI
	o.Code = code

	return o
}

func (o *OAuth) ToJSON() string {
	json, _ := json.Marshal(o)
	return string(json)
}
