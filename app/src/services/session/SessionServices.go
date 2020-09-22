package session

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yoshi0202/grass-app/app/src/constant"
	"github.com/yoshi0202/grass-app/app/src/dto"
	"github.com/yoshi0202/grass-app/app/src/services"
)

func FindByGithub(token string) []byte {
	to := constant.NewConst().GetUserAPIURL
	req, _ := http.NewRequest("GET", to, nil)
	req.Header.Set("Authorization", token)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray
}

func GetAccessToken(code string) string {
	to := constant.NewConst().AccessTokenURL
	oauth := dto.NewOAuth(code)
	json, _ := json.Marshal(oauth)
	resp, _ := http.Post(to, "application/json", bytes.NewBuffer(json))
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	accessToken := services.CreateRegexp(`[&=]`).Split(string(byteArray), -1)[1]
	return accessToken
}
