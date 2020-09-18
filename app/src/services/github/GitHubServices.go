package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yoshi0202/grass-app/app/src/constant"
)

type GithubGrass struct {
	Count string `json:"count"`
	Date  string `json:"date"`
}

type GithubGrasses []GithubGrass

func GetGrass(param string) []byte {
	c := constant.NewConst()
	url := c.Url + param + "/contributions"
	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray
}

func CreateGrasses(array GithubGrasses, apiRes [][]byte) string {
	for _, v := range apiRes {
		g := GithubGrass{}
		g.Count = strings.Split(strings.Split(string(v), " ")[7], "\"")[1]
		g.Date = strings.Split(strings.Split(string(v), " ")[8], "\"")[1]
		array = append(array, g)
	}
	result, _ := json.Marshal(array)
	return string(result)
}
