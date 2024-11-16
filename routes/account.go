package routes

import (
	"github/teacoder-team/golang-backend/services"
	"github/teacoder-team/golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserDTO представляет данные, необходимые для создания пользователя
// @Description Данные, используемые для создания нового пользователя
type UserDTO struct {
	Email       string `json:"email" binding:"required" example:"user@example.com"`
	Password    string `json:"password" binding:"required" example:"password123"`
	DisplayName string `json:"display_name" binding:"required" example:"Иван Иванов"`
}

// CreateUserHandler Обработчик запроса для создания нового пользователя
// @Summary Создать нового пользователя
// @Description Данный метод создает нового пользователя на основе предоставленных данных.
// @Tags Account
// @Accept json
// @Produce json
// @Param user body UserDTO true "Данные для создания нового пользователя"
// @Success 200 {object} UserDTO "Пользователь успешно создан"
// @Failure 400 {object} utils.ErrorResponse "Некорректные данные для создания пользователя"
// @Failure 409 {object} utils.ErrorResponse "Пользователь с таким email уже существует"
// @Router /auth/account/create [post]
func CreateUserHandler(c *gin.Context) {
	var userDto UserDTO

	if err := c.ShouldBindJSON(&userDto); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := services.CreateUser(userDto.Email, userDto.Password, userDto.DisplayName, c)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func RegisterAccountRoutes(r *gin.Engine) {
	r.POST("/auth/account/create", CreateUserHandler)
}
