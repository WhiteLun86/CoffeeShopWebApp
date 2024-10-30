package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// 创建一个新的 Echo 实例
	e := echo.New()
	// 定义一个路由
	e.GET("/", helloHandler)

	// 启动服务器
	e.Logger.Fatal(e.Start(":1323"))
}
