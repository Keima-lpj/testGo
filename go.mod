module testGo

go 1.16

require (
	gioui.org v0.0.0-20220601100144-a896a467ecae
	github.com/LiaoPuJian/keima v0.0.0-20210909124453-bfbe28e8e314
	github.com/OpenWhiteBox/AES v0.0.0-20161114232003-b7fcb3c27b63
	github.com/OpenWhiteBox/primitives v0.0.0-20161020045608-2f25eea09f86
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/astaxie/beego v1.12.2
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/garyburd/redigo v1.6.0
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/godbus/dbus/v5 v5.1.0
	github.com/golang/protobuf v1.5.2
	github.com/google/wire v0.4.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/gorilla/websocket v1.4.2
	github.com/henrylee2cn/faygo v1.3.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic v6.2.34+incompatible
	github.com/panjf2000/ants v1.3.0
	github.com/panjf2000/ants/v2 v2.5.0
	github.com/pyroscope-io/client v0.2.1
	github.com/satori/go.uuid v1.2.0
	github.com/xuyu/goredis v0.0.0-20160929021245-89fbe9474b37
	github.com/zeromicro/go-zero v1.3.3
	golang.org/x/exp v0.0.0-20210722180016-6781d3edade3
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
	pkg.deepin.com/golang/lib/testing v0.0.0-20210825024834-7b5697905657
	pkg.deepin.com/golang/lib/uuid v0.0.0-20200417051548-1fe45320d2ca
	pkg.deepin.com/web/unionid/util v0.0.0-20211012074024-ebf44e0908d3
)

require (
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.2.4 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/mailru/easyjson v0.7.2 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/temoto/robotstxt v1.1.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.0 // indirect
)

replace (
	github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.6.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
