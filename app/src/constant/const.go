package constant

type Constant struct {
	SvgTag     string
	RectTag    string
	Url        string
	DbUser     string
	DbPassword string
	DbProtocol string
	DbName     string
}

func NewConst() *Constant {
	c := new(Constant)
	c.SvgTag = `<svg(?: [\s\S]+?)?>[\s\S]*?<\/svg>`
	c.RectTag = `<rect.*?\/>`
	c.Url = "https://github.com/users/"
	c.DbUser = "root"
	c.DbPassword = "password"
	c.DbProtocol = "tcp(db:3306)"
	c.DbName = "grass_app"

	return c
}
