package services

import (
	"github/teacoder-team/golang-backend/config"
	"github/teacoder-team/golang-backend/models"
	"github/teacoder-team/golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(email, password, displayName string, c *gin.Context) (*models.User, error) {
	var existingUser models.User

	if err := config.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		utils.SendErrorResponse(c, http.StatusConflict, "Такой пользователь уже существует")
		return nil, err
	}

	hashedPassword, err := utils.HashPasswordToBase64(password)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Ошибка при хэшировании пароля")
		return nil, err
	}

	user := &models.User{
		ID:          uuid.New().String(),
		Email:       email,
		Password:    &hashedPassword,
		DisplayName: displayName,
		Username:    "teacoder",
		Method:      models.MethodCreds,
		Role:        models.RoleStudent,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Ошибка при создании пользователя")
		return nil, err
	}

	return user, nil
}
