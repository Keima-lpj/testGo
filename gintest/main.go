package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

var sessionName = "gin-session"
var sessionSecureCodec []securecookie.Codec

func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	sessionSecureCodec = securecookie.CodecsFromPairs([]byte(secret))
	// Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: false, Path: "/"})
	return sessions.Sessions(sessionName, store)
}

func TestCookie(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Set("user", "111")
	sess.Save()
}

func GetUser(c *gin.Context) {
	sess := sessions.Default(c)
	sessionID, _ := c.Cookie(sessionName)
	fmt.Println("sessionID", sessionID)
	fmt.Println("aaa", sess.Get("user"))
}

func main() {
	g := gin.Default()
	g.Use(Session("123"))
	g.GET("/a", TestCookie)
	g.GET("/user", GetUser)
	g.Run(":9999")
}
