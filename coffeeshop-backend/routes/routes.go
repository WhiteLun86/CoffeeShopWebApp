// routes/routes.go
package routes

import (
	"gobackend/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/login", controllers.Login)
	e.GET("/api/users", controllers.GetUsers)
	e.DELETE("api/users/:id", controllers.DeleteUser) // 添加删除用户路由
}
