package controllers

import (
	"gobackend/config"
	"gobackend/models"
	"net/http"

	"log"

	"github.com/labstack/echo/v4"
)

type OrderRequest struct {
	UserID     uint                  `json:"user_id"`
	Products   []OrderProductRequest `json:"products"`    // 包含多个商品请求
	TotalPrice float64               `json:"total_price"` // 总价格
}

type OrderProductRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func CreateOrder(c echo.Context) error {
	req := new(OrderRequest)
	if err := c.Bind(req); err != nil {
		log.Printf("请求绑定错误: %v", err) // 调试信息
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	// 检查库存
	for _, productReq := range req.Products {
		var coffeeProduct models.CoffeeProduct
		if err := config.DB.First(&coffeeProduct, productReq.ProductID).Error; err != nil {
			log.Printf("获取产品失败: %v", err) // 调试信息
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "商品不存在"})
		}

		if coffeeProduct.Stock < productReq.Quantity {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "库存不足"})
		}
	}

	order := models.Order{
		UserID:     req.UserID,
		TotalPrice: req.TotalPrice,
		Status:     "pending", // 默认状态
	}

	// 创建订单
	if err := config.DB.Create(&order).Error; err != nil {
		log.Printf("创建订单失败: %v", err) // 调试信息
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "创建订单失败"})
	}

	// 插入产品到 order_products 表
	for _, productReq := range req.Products {
		orderProduct := models.OrderProduct{
			OrderID:   order.ID,
			ProductID: productReq.ProductID,
			Quantity:  productReq.Quantity,
			Price:     0, // 价格将在下方获取
		}

		// 获取产品信息以设置价格
		var coffeeProduct models.CoffeeProduct
		if err := config.DB.First(&coffeeProduct, orderProduct.ProductID).Error; err == nil {
			orderProduct.Price = coffeeProduct.Price // 从产品表中获取价格
		}

		if err := config.DB.Create(&orderProduct).Error; err != nil {
			log.Printf("添加产品到订单失败: %v", err) // 调试信息
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "添加产品到订单失败"})
		}

		// 更新库存
		coffeeProduct.Stock -= productReq.Quantity
		if err := config.DB.Save(&coffeeProduct).Error; err != nil {
			log.Printf("更新库存失败: %v", err) // 调试信息
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "更新库存失败"})
		}
	}

	return c.JSON(http.StatusCreated, order)
}

func GetOrders(c echo.Context) error {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "获取订单失败"})
	}

	// 遍历订单并获取每个订单的产品信息
	for i := range orders {
		var orderProducts []models.OrderProduct

		if err := config.DB.Where("order_id = ?", orders[i].ID).
			Find(&orderProducts).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "获取产品失败"})
		}

		// 将产品信息添加到订单的 Products 字段
		for _, op := range orderProducts {
			var coffeeProduct models.CoffeeProduct
			if err := config.DB.First(&coffeeProduct, op.ProductID).Error; err == nil {
				op.Product = coffeeProduct // 获取完整的产品信息
			}
			orders[i].Products = append(orders[i].Products, op)
		}
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

func GetUserOrders(c echo.Context) error {
	// 从 URL 参数中获取用户 ID
	userId := c.Param("user_id") // 假设在路由中定义了 `/orders/:user_id`

	var orders []models.Order
	// 查询特定用户的订单
	if err := config.DB.Where("user_id = ?", userId).Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "获取订单失败"})
	}

	// 遍历订单并获取每个订单的产品信息
	for i := range orders {
		var orderProducts []models.OrderProduct

		if err := config.DB.Where("order_id = ?", orders[i].ID).
			Find(&orderProducts).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "获取产品失败"})
		}

		// 将产品信息添加到订单的 Products 字段
		for _, op := range orderProducts {
			var coffeeProduct models.CoffeeProduct
			if err := config.DB.First(&coffeeProduct, op.ProductID).Error; err == nil {
				op.Product = coffeeProduct // 获取完整的产品信息
			}
			orders[i].Products = append(orders[i].Products, op)
		}
	}

	return c.JSON(http.StatusOK, orders)
}
