package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshi0202/grass-app/app/src/constant"
	"github.com/yoshi0202/grass-app/app/src/services"
	"github.com/yoshi0202/grass-app/app/src/services/github"
	"github.com/yoshi0202/grass-app/app/src/services/users"
)

func main() {
	router := gin.Default()

	// root
	router.GET("/", func(c *gin.Context) {
		user := users.FindAllUser()
		fmt.Println(user)
		c.String(http.StatusOK, "Ok")
	})

	router.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "login page")
	})

	router.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "create session")
	})

	router.DELETE("/logout", func(c *gin.Context) {
		c.String(http.StatusOK, "logout")
	})

	router.GET("/signup", func(c *gin.Context) {
		c.String(http.StatusOK, "signup page")
	})

	router.POST("/signup", func(c *gin.Context) {
		c.String(http.StatusOK, "signup")
	})

	router.GET("/timeline", func(c *gin.Context) {
		c.String(http.StatusOK, "timeline page")
	})

	router.GET("/regist", func(c *gin.Context) {
		c.String(http.StatusOK, "enter github id pages")
	})

	router.POST("/regist", func(c *gin.Context) {
		c.String(http.StatusOK, "regist user")
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		con := constant.NewConst()
		apiRes := github.GetGrass(c.Param("name"))
		svgTags := services.FindSvgTag(con.SvgTag, apiRes)
		rectArray := services.FindAllRectTag(con.RectTag, svgTags)
		grasses := github.CreateGrasses(github.GithubGrasses{}, rectArray)
		c.String(http.StatusOK, grasses)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	// router.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	message := name + " is " + action
	// 	c.String(http.StatusOK, message)
	// })

	// For each matched request Context will hold the route definition
	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	c.FullPath() == "/user/:name/*action" // true
	// })

	router.Run(":8080")
}
