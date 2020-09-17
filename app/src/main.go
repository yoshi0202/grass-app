package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yoshi0202/grass-app/app/src/services"
)

type GithubGrass struct {
	Count string `json:"count"`
	Date  string `json:"date"`
}

type GithubGrasses []GithubGrass

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		apiRes := services.Get(c.Param("name"))
		svgTags := services.FindSvgTag(`<svg(?: [\s\S]+?)?>[\s\S]*?<\/svg>`, apiRes)
		rectArray := services.FindAllRectTag(`<rect.*?\/>`, svgTags)
		gl := GithubGrasses{}
		for _, v := range rectArray {
			g := GithubGrass{}
			g.Count = strings.Split(strings.Split(string(v), " ")[7], "\"")[1]
			g.Date = strings.Split(strings.Split(string(v), " ")[8], "\"")[1]
			gl = append(gl, g)
		}
		result, _ := json.Marshal(gl)

		c.String(http.StatusOK, string(result))
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	c.FullPath() == "/user/:name/*action" // true
	// })

	router.Run(":8080")
}
