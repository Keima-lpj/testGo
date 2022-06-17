package main

import (
	"github.com/henrylee2cn/faygo"
	//"time"
	"time"
)

/**
 * index model
 */
type Index struct {
	Id        int      `param:"<in:path> <required> <desc:ID> <range: 0:10>"`
	Title     string   `param:"<in:query> <nonzero>"`
	Paragraph []string `param:"<in:query> <name:p> <len: 1:10> <regexp: ^[\\w]*$>"`
	Cookie    string   `param:"<in:cookie> <name:faygoID>"`
	// Picture         *multipart.FileHeader `param:"<in:formData> <name:pic> <maxmb:30>"`
}

func (i *Index) Serve(ctx *faygo.Context) error {
	if ctx.CookieParam("faygoID") == "" {
		ctx.SetCookie("faygoID", time.Now().String())
	}
	return ctx.JSON(200, i)
}

/**
 * 测试faygo框架
 */
func main() {
	//生成新的框架
	app := faygo.New("myapp", "1.0")
	//注册路由
	app.GET("/index/:id", new(Index))
	//框架启动
	faygo.Run()
}
