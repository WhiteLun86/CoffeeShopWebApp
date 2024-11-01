package controllers

import (
	"gobackend/config"
	"gobackend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 新的 UserResponse 结构体，包含原 User 结构体的所有字段
type UserResponse struct {
	Key      int    `json:"key"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetUsers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "获取用户数据失败"})
	}

	var userResponses []UserResponse
	for i, user := range users {
		userResponses = append(userResponses, UserResponse{
			Key:      i,             // 从 0 开始的 key
			ID:       user.ID,       // User 的 ID
			Username: user.Username, // User 的 Username
			Password: user.Password, // User 的 Password
			Role:     user.Role,     // User 的 Role
		})
	}

	return c.JSON(http.StatusOK, userResponses)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id") // 从 URL 获取用户 ID

	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "用户不存在"})
	}

	// 删除用户
	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "删除用户失败"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "用户删除成功"})
}
