package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type GithubGrass struct {
	count string `json:"count"`
	date  string `json:"date"`
}
type GithubGrasses []GithubGrass

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		// name := c.Param("name")
		url := "https://github.com/users/yoshi0202/contributions"
		req, _ := http.NewRequest("GET", url, nil)

		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		re := regexp.MustCompile(`<svg(?: [\s\S]+?)?>[\s\S]*?<\/svg>`)
		svg := re.Find([]byte(string(byteArray)))
		re2 := regexp.MustCompile(`<rect.*?\/>`)
		rect := re2.FindAll([]byte(string(svg)), -1)
		// var arr []GithubGrass
		gl := GithubGrasses{}
		for _, v := range rect {
			g := GithubGrass{}
			g.count = strings.Split(strings.Split(string(v), " ")[7], "\"")[1]
			g.date = strings.Split(strings.Split(string(v), " ")[8], "\"")[1]
			gl = append(gl, g)
		}
		fmt.Println(gl)
		result, _ := json.Marshal(gl)
		fmt.Printf("[+] %s\n", string(result))

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
