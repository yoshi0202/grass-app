package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/yoshi0202/grass-app/app/src/dto"
	"github.com/yoshi0202/grass-app/app/src/services"
	"github.com/yoshi0202/grass-app/app/src/services/grass"
	"github.com/yoshi0202/grass-app/app/src/services/session"
	"github.com/yoshi0202/grass-app/app/src/services/usergrassess"
	"github.com/yoshi0202/grass-app/app/src/services/users"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("*.tmpl")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// root
	router.GET("/", func(c *gin.Context) {
		user := users.FindAll()
		c.String(http.StatusOK, string(user.ToJSON()))
	})

	router.GET("/migrate", func(c *gin.Context) {
		user := dto.NewUser()
		grass := dto.NewGrass()
		userGrass := new(usergrassess.UserGrassess)
		db := services.ConnectGorm()
		db.AutoMigrate(user, grass, userGrass)
		c.String(http.StatusOK, "migrate ok")
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	})

	// router.DELETE("/logout", checkCookie(), func(c *gin.Context) {
	router.GET("/logout", checkCookie(), func(c *gin.Context) {
		session.DeleteSession(c)
		c.Redirect(302, "/")
	})

	router.GET("/timeline", checkCookie(), func(c *gin.Context) {
		c.String(http.StatusOK, "timeline page")
	})

	router.GET("/github/callback", func(c *gin.Context) {
		token := session.GetAccessToken(c.Query("code"))
		session.FindByGithub("token " + token)
		session.CreateLoginSession(token, c)
		c.Redirect(302, "/timeline")
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		g := grass.FindByGithub(c.Param("name"))
		grass.FindOrCreate(c.Param("name"), g)
		c.String(http.StatusOK, g.ToJSON())
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

func checkCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		userID := sess.Get("userId")
		if userID == nil {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Set("userId", userID)
		c.Next()
	}
}
