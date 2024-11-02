// controllers/order.go
package controllers

import (
	"gobackend/config"
	"gobackend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderRequest struct {
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func CreateOrder(c echo.Context) error {
	req := new(OrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	order := models.Order{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Price:     req.Price,
		Status:    "pending", // 默认状态
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "创建订单失败"})
	}

	return c.JSON(http.StatusCreated, order)
}

func GetOrders(c echo.Context) error {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "获取订单失败"})
	}
	return c.JSON(http.StatusOK, orders)
}

func DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Order{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "订单不存在"})
	}
	return c.NoContent(http.StatusNoContent)
}

// controllers/order.go
func UpdateOrderStatus(c echo.Context) error {
	id := c.Param("id")
	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "订单未找到"})
	}

	// 更新订单状态
	order.Status = req.Status
	if err := config.DB.Save(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "更新订单状态失败"})
	}

	return c.JSON(http.StatusOK, order)
}
