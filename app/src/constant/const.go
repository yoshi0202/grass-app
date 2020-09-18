package constant

type Constant struct {
	SvgTag  string
	RectTag string
	Url     string
}

func NewConst() *Constant {
	c := new(Constant)
	c.SvgTag = `<svg(?: [\s\S]+?)?>[\s\S]*?<\/svg>`
	c.RectTag = `<rect.*?\/>`
	c.Url = "https://github.com/users/"

	return c
}
