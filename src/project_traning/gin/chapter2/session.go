package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

var sessionName string
var sessionValue string

type MyOption struct {
	sessions.Options
}

// go get github.com/gin-contrib/sessions
func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("session_secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/session", func(context *gin.Context) {
		name := context.Query("name")
		if len(name) <= 0 {
			context.JSON(http.StatusBadRequest, "data wrong")
			return
		}
		sessionName = "session_" + name
		sessionValue = "session_value_" + name
		session := sessions.Default(context)
		sessionData := session.Get(sessionName)
		if sessionData != sessionValue {
			session.Set(sessionName, sessionValue)
			o := MyOption{}
			o.Path = "/"
			o.MaxAge = 10
			session.Options(o.Options)
			session.Save()
			context.JSON(http.StatusOK, "session has saved")
			return
		}
		context.JSON(http.StatusOK, "your session is "+sessionData.(string))

	})
	r.Run(":9090")

}
