package constant

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Constant struct {
	SvgTag         string
	RectTag        string
	Url            string
	DbUser         string
	DbPassword     string
	DbProtocol     string
	DbName         string
	AccessTokenURL string
	GetUserAPIURL  string
	ClientID       string
	ClientSecret   string
	RedirectURI    string
}

func NewConst() *Constant {
	godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	c := &Constant{}
	c.SvgTag = `<svg(?: [\s\S]+?)?>[\s\S]*?<\/svg>`
	c.RectTag = `<rect.*?\/>`
	c.Url = "https://github.com/users/"
	c.DbUser = "root"
	c.DbPassword = "password"
	c.DbProtocol = "tcp(db:3306)"
	c.DbName = "grass_app"
	c.AccessTokenURL = os.Getenv("ACCESS_TOKEN_URL")
	c.GetUserAPIURL = os.Getenv("GET_USER_API_URL")
	c.ClientID = os.Getenv("CLIENT_ID")
	c.ClientSecret = os.Getenv("CLIENT_SECRET")
	c.RedirectURI = os.Getenv("REDIRECT_URI")
	return c
}
