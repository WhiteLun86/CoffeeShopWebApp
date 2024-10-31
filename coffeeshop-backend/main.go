// main.go
package main

import (
	"gobackend/config"
	"gobackend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// 连接数据库
	config.ConnectDatabase()

	// 创建 Echo 实例
	e := echo.New()

	// 初始化路由
	routes.InitRoutes(e)

	// 启动服务器
	e.Logger.Fatal(e.Start(":8080"))
}
