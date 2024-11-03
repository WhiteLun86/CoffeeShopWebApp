// routes/routes.go
package routes

import (
	"gobackend/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/login", controllers.Login)

	e.GET("/api/users", controllers.GetUsers)
	e.DELETE("api/users/:id", controllers.DeleteUser)
	e.PUT("/api/user/update", controllers.UpdateUser)
	e.POST("/api/users", controllers.CreateUser)

	e.GET("/api/products", controllers.GetCoffeeProducts)
	e.DELETE("/api/products/:id", controllers.DeleteCoffeeProduct)
	e.POST("/api/products", controllers.AddCoffeeProduct)

	e.POST("/api/orders", controllers.CreateOrder)
	e.GET("/api/orders", controllers.GetOrders)
	e.DELETE("/api/orders/:id", controllers.DeleteOrder)
	e.PUT("/api/coffee/:id", controllers.UpdateProduct)
	e.PUT("/api/orders/:id/status", controllers.UpdateOrderStatus)
	e.GET("/api/user/:user_id/orders", controllers.GetUserOrders)
}
