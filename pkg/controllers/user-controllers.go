package controllers

import (
	"errors"
	"net/http"

	"caching-user-app/models"
	"caching-user-app/pkg/database"
	"caching-user-app/pkg/transformator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUsers(ctx *gin.Context) {
	var users []models.User

	db := database.GetDatabaseConnection()
	result := db.Find(&users)

	if result.Error != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": result.Error.Error()},
		)
		return
	}

	usersResponse := make([]models.UserResponse, len(users))

	for index, user := range users {
		usersResponse[index] = transformator.ToUserResponse(user)
	}

	ctx.JSON(http.StatusOK, usersResponse)
}

func GetUser(ctx *gin.Context) {
	var user models.User

	id := ctx.Param("id")
	db := database.GetDatabaseConnection()
	result := db.First(&user, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userResponse := transformator.ToUserResponse(user)

	if result.Error == nil {
		ctx.JSON(http.StatusOK, userResponse)
		return
	}

	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"error": result.Error.Error()},
	)
}

func CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDatabaseConnection()
	user.ID = uuid.New().String()
	result := db.Create(&user)

	if result.Error != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": result.Error.Error()},
		)
		return
	}

	userResponse := transformator.ToUserResponse(user)

	ctx.JSON(
		http.StatusCreated,
		gin.H{"message": "User created successfully", "user": userResponse},
	)
}
