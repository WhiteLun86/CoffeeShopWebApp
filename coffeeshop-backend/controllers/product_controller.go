package controllers

import (
	"gobackend/config"
	"gobackend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCoffeeProducts(c echo.Context) error {
	var products []models.CoffeeProduct
	if err := config.DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "获取商品数据失败"})
	}
	return c.JSON(http.StatusOK, products)
}

func DeleteCoffeeProduct(c echo.Context) error {
	id := c.Param("id")

	// 从数据库中查找商品
	var product models.CoffeeProduct
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "商品不存在"})
	}

	// 删除商品
	if err := config.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "删除失败"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "商品已删除"})
}

func AddCoffeeProduct(c echo.Context) error {
	var coffeeProduct models.CoffeeProduct

	// 解析请求体
	if err := c.Bind(&coffeeProduct); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	// 将商品添加到数据库
	if err := config.DB.Create(&coffeeProduct).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "添加商品失败"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "商品添加成功", "product": coffeeProduct})
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

func UpdateProduct(c echo.Context) error {
	var req UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	id := c.Param("id")

	var product models.CoffeeProduct
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "商品未找到"})
	}

	// 更新商品信息
	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock
	product.Description = req.Description

	if err := config.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "更新商品失败"})
	}

	return c.JSON(http.StatusOK, product)
}
