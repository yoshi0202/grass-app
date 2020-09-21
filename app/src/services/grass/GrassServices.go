package grass

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/yoshi0202/grass-app/app/src/constant"
	"github.com/yoshi0202/grass-app/app/src/services"
)

type Grass struct {
	gorm.Model
	Id        string         `json:"id"`
	GitHubID  string         `json:"githubId"`
	Count     string         `json:"count"`
	Date      string         `json:"date"`
	CountDate datatypes.JSON `json:"countDate"`
}

type Grasses []Grass

func Find(param int) *Grass {
	g := new(Grass)
	db := services.ConnectGorm()
	db.First(&g, param)
	return g
}

func FindAll() *Grasses {
	g := new(Grasses)
	db := services.ConnectGorm()
	db.Find(&g)
	return g
}

func FindOrCreate(githubId string, json []byte) string {
	if len(json) == 0 {
		return ""
	}
	g := new(Grass)
	db := services.ConnectGorm()
	g.GitHubID = githubId
	g.CountDate = datatypes.JSON([]byte(json))
	db.FirstOrCreate(&g, Grass{GitHubID: githubId})
	return "ok"
}

func FindByGithub(name string) []byte {
	con := constant.NewConst()
	apiRes := GetGrass(name)
	svgTags := services.FindSvgTag(con.SvgTag, apiRes)
	rectArray := services.FindAllRectTag(con.RectTag, svgTags)
	grasses := CreateGrasses(Grasses{}, rectArray)
	return grasses
}

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

func CreateGrasses(array Grasses, apiRes [][]byte) []byte {
	for _, v := range apiRes {
		g := Grass{}
		g.Count = strings.Split(strings.Split(string(v), " ")[7], "\"")[1]
		g.Date = strings.Split(strings.Split(string(v), " ")[8], "\"")[1]
		array = append(array, g)
	}
	result, _ := json.Marshal(array)
	return result
}
