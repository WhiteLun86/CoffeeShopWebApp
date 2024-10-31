// controllers/auth.go
package controllers

import (
	"gobackend/config"
	"gobackend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		// 返回404状态码和错误信息，明确用户不存在
		return c.JSON(http.StatusNotFound, echo.Map{"message": "用户不存在"})
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		// 返回401状态码和错误信息，提示密码错误
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "密码错误"})
	}

	// 返回用户角色和登录成功信息
	return c.JSON(http.StatusOK, echo.Map{
		"message": "登录成功",
		"role":    user.Role,
	})
}
