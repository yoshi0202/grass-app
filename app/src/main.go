package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// root
	router.GET("/", func(c *gin.Context) {
		user := users.FindAll()
		c.String(http.StatusOK, string(user.ToJSON()))
	})

	router.GET("/owngrass", checkCookie(), func(c *gin.Context) {
		u, _ := c.Get("userId")
		grass := grass.FindByGitHubID(u.(string))
		c.String(http.StatusOK, grass.ToJSON())
	})

	router.GET("/migrate", func(c *gin.Context) {
		user := dto.NewUser()
		grass := dto.NewGrass()
		userGrass := dto.NewUserGrass()
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
		c.Redirect(302, "http://localhost:3000/")
	})

	router.GET("/timeline", checkCookie(), func(c *gin.Context) {
		u, _ := c.Get("userId")
		ug := usergrassess.FindAllByUserId(u.(string))
		g := grass.SearchByUserGrasses(ug)
		c.String(http.StatusOK, g.ToJSON())
	})

	router.PUT("/timeline/delete", checkCookie(), func(c *gin.Context) {
		ug := dto.NewUserGrass()
		c.Bind(ug)
		ui, _ := c.Get("userId")
		ug.LoginID = ui.(string)
		usergrassess.Delete(ug)
		c.String(http.StatusOK, "delete")
	})

	router.POST("/timeline/regist", checkCookie(), func(c *gin.Context) {
		ut := dto.NewUserTimeline()
		c.Bind(ut)
		u, _ := c.Get("userId")
		g := grass.FindByGithub(ut.GitHubID)
		gr := grass.FindOrCreate(ut.GitHubID, g)
		if len(gr.CountDate) < 3 {
			c.String(http.StatusOK, "user not found")
		} else {
			usergrassess.Create(u.(string), gr)
			c.String(http.StatusOK, "success!")
		}
	})

	router.GET("/github/callback", func(c *gin.Context) {
		token := session.GetAccessToken(c.Query("code"))
		login := session.FindByGithub("token " + token)
		oauthUser := users.GetUserIDByJSON(login)

		// ユーザが存在していなければ作成
		users.FindOrCreate(oauthUser)
		g := grass.FindByGithub(oauthUser.Login)
		grass.FindOrCreate(oauthUser.Login, g)
		session.CreateLoginSession(oauthUser.Login, c)
		// c.String(http.StatusOK, user.ToJSON())
		c.Redirect(302, "http://localhost:3000/timeline")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		g := grass.FindByGithub(c.Param("name"))
		grass.FindOrCreate(c.Param("name"), g)
		c.String(http.StatusOK, g.ToJSON())
	})

	router.Run(":8080")

}

// middleware
func checkCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		userID := sess.Get("userId")
		if userID == nil {
			// c.Redirect(302, "http://localhost:3000/login")
			c.String(http.StatusOK, "redirect")
			c.Abort()
			return
		}
		c.Set("userId", userID)
		c.Next()
	}
}
