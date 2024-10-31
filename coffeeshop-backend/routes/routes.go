// routes/routes.go
package routes

import (
	"gobackend/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/login", controllers.Login)
}
