package constant

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
	c := &Constant{}
	c.SvgTag = `<svg(?: [\s\S]+?)?>[\s\S]*?<\/svg>`
	c.RectTag = `<rect.*?\/>`
	c.Url = "https://github.com/users/"
	c.DbUser = "root"
	c.DbPassword = "password"
	c.DbProtocol = "tcp(db:3306)"
	c.DbName = "grass_app"
	c.AccessTokenURL = "https://github.com/login/oauth/access_token"
	c.GetUserAPIURL = "https://api.github.com/user"
	c.ClientID = "6275811b400983aaba52"
	c.ClientSecret = "a9a4d5932feb8b0261912e481ab862a979922a2a"
	c.RedirectURI = "http://localhost:8080/github/callback"
	return c
}
