package main

import (
	"encoding/base32"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	gs "github.com/gorilla/sessions"
	"pkg.deepin.com/golang/lib/uuid"
)

type innerSession interface {
	Session() *gs.Session
}

var sessionSecureCodec []securecookie.Codec

func main() {
	g := gin.Default()

	sessionSecureCodec = securecookie.CodecsFromPairs([]byte("secret"))
	store, err := redis.NewStoreWithDB(10, "tcp", "redis-master.sndu.cn:6379", "deepin!@#", "22", []byte("secret"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   1200,
		HttpOnly: true,
	})
	g.Use(sessions.Sessions("uniontech1", store))
	g.Use(func(c *gin.Context) {
		c.Next()
		s := sessions.Default(c)
		gsi := s.(innerSession).Session()
		if gsi.ID == "" {
			gsi.ID = strings.TrimRight(base32.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32)), "=")
		}
		encoded, _ := securecookie.EncodeMulti(gsi.Name(), gsi.ID, sessionSecureCodec...)
		option := *gsi.Options
		option.MaxAge = 0
		http.SetCookie(c.Writer, gs.NewCookie(gsi.Name(), encoded, &option))
	})

	g.Any("set", func(ctx *gin.Context) {
		s := sessions.Default(ctx)
		id := uuid.UUID32()
		s.Set("key", id)
		s.Save()
		ctx.JSON(200, id)
	})

	g.Any("get", func(ctx *gin.Context) {
		s := sessions.Default(ctx)
		id := s.Get("key")
		ctx.JSON(200, id)
	})

	g.Run()
}
