package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

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
		for _, v := range rect {
			fmt.Println(string(v))
		}
		c.String(http.StatusOK, string(svg))
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
