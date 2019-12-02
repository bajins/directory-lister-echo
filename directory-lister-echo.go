package main

import (
	"flag"
	"github.com/labstack/echo/v4"
)

// 禁止浏览器页面缓存
func FilterNoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("Expires", "0")
		return next(c)
	}
}

// 处理跨域请求,支持options访问
func Cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := c.Request().Method
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			return next(c)
		}
		// 它指定允许进入来源的域名、ip+端口号 。 如果值是 ‘*’ ，表示接受任意的域名请求，这个方式不推荐，
		// 主要是因为其不安全，而且因为如果浏览器的请求携带了cookie信息，会发生错误
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		// 设置服务器允许浏览器发送请求都携带cookie
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		// 允许的访问方法
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE, PATCH")
		// Access-Control-Max-Age 用于 CORS 相关配置的缓存
		c.Response().Header().Set("Access-Control-Max-Age", "3600")
		// 设置允许的请求头信息
		c.Response().Header().Set("Access-Control-Allow-Headers", "Token,Origin, X-Requested-With, Content-Type, Accept,mid,X-Token,AccessToken,X-CSRF-Token, Authorization")

		c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		return next(c)
	}
}

// 获取传入参数的端口，如果没传默认值为8000
func Port() (port string) {
	flag.StringVar(&port, "p", "8000", "默认端口:8000")
	flag.Parse()
	return ":" + port

	//if len(os.Args[1:]) == 0 {
	//	return ":8000"
	//}
	//return ":" + os.Args[1]
}

func main() {
	e := echo.New()
	e.Use(FilterNoCache)
	//e.Use(Cors())
	//e.Use(Authorize())

	e.Static("static", "static")
	e.GET("/dir", GetDir)
	e.Any("/", GetDir)
	e.Any("/*", GetDir)
	e.Any("/admin/*", Admin)
	e.Logger.Fatal(e.Start(Port()))
}
