package controllers

import (
	"gobackend/config"
	"gobackend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

// 更新用户信息处理函数
func UpdateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	// 查找要更新的用户
	var existingUser models.User
	if err := config.DB.First(&existingUser, user.ID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "用户未找到"})
	}

	// 更新字段（包括密码加密）
	existingUser.Username = user.Username
	existingUser.Role = user.Role
	if user.Password != "" { // 仅在用户输入新密码时更新
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		existingUser.Password = string(hashedPassword)
	}

	// 保存更改
	if err := config.DB.Save(&existingUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "更新用户失败"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "用户更新成功"})
}

// CreateUserRequest 用于接收前端请求
type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=user admin"`
}

// CreateUser 处理创建用户请求
func CreateUser(c echo.Context) error {
	req := new(CreateUserRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "请求格式错误"})
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "密码加密失败"})
	}

	// 创建用户
	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "创建用户失败"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "用户创建成功"})
}
