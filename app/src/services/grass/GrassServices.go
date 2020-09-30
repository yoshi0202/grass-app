package grass

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yoshi0202/grass-app/app/src/constant"
	"github.com/yoshi0202/grass-app/app/src/dto"
	"github.com/yoshi0202/grass-app/app/src/services"
)

func Find(param int) *dto.Grass {
	g := dto.NewGrass()
	db := services.ConnectGorm()
	db.First(&g, param)
	return g
}

func Update(g *dto.Grass) *dto.Grass {
	db := services.ConnectGorm()
	db.Model(&dto.Grass{}).Where("git_hub_id = ?", g.GitHubID).Update("count_date", g.CountDate)
	return g
}

func FindByGitHubID(param string) *dto.Grass {
	g := dto.NewGrass()
	db := services.ConnectGorm()
	db.Where("git_hub_id = ?", param).Find(&g)
	return g
}

func FindAll() *dto.Grasses {
	g := dto.NewGrasses()
	db := services.ConnectGorm()
	db.Find(&g)
	return g
}

func SearchByUserGrasses(ug *dto.UserGrasses) *dto.Grasses {
	var arr []uint
	for _, v := range *ug {
		arr = append(arr, v.GrassID)
	}
	g := dto.NewGrasses()
	db := services.ConnectGorm()
	db.Where("id IN ?", arr).Find(&g)
	return g
}

func FindOrCreate(githubId string, grass *dto.Grass) *dto.Grass {
	if len(grass.CountDate) < 3 {
		return grass
	}
	db := services.ConnectGorm()
	db.FirstOrCreate(&grass, dto.Grass{GitHubID: githubId})
	return grass
}

func FindByGithub(name string) *dto.Grass {
	gs := dto.NewGrass()
	gs.GitHubID = name
	con := constant.NewConst()
	apiRes := GetGrass(name)
	svgTags := services.FindSvgTag(con.SvgTag, apiRes)
	rectArray := services.FindAllRectTag(con.RectTag, svgTags)
	grass := CreateGrasses(gs, rectArray, name)
	return grass
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

func CreateGrasses(gs *dto.Grass, apiRes [][]byte, name string) *dto.Grass {
	gitHubs := dto.NewGitHubs()
	fmt.Println(len(apiRes))
	for _, v := range apiRes {
		g := dto.NewGitHub()
		g.Count = strings.Split(strings.Split(string(v), " ")[7], "\"")[1]
		g.Date = strings.Split(strings.Split(string(v), " ")[8], "\"")[1]
		*gitHubs = append([]dto.GitHub{*g}, *gitHubs...)
	}
	gs.CountDate = gitHubs.ToJSON()
	return gs
}

func UpdateGrasses() {
	grasses := FindAll()
	for _, v := range *grasses {
		newGrass := FindByGithub(v.GitHubID)
		Update(newGrass)
	}
}
