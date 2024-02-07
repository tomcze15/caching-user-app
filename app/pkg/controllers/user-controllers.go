package controllers

import (
	"errors"
	"net/http"

	"caching-user-app/globals"
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

func tryResponseWithCachedUser(ctx *gin.Context, id string) bool {
	cachedUser, isFound := globals.Cache.Get(id)
	if isFound {
		responseWithUser(ctx, cachedUser, true)
		return true
	}
	return false
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if tryResponseWithCachedUser(ctx, id) {
		return
	}

	lock := globals.Cache.LockForKey(id)
	lock.Lock()
	defer lock.Unlock()

	if tryResponseWithCachedUser(ctx, id) {
		return
	}

	db := database.GetDatabaseConnection()
	var user models.User

	result := db.First(&user, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if result.Error != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": result.Error.Error()},
		)
		return
	}

	userResponse := transformator.ToUserResponse(user)
	globals.Cache.Set(id, userResponse)

	responseWithUser(ctx, userResponse, false)
}

func responseWithUser(
	ctx *gin.Context,
	userResponse models.UserResponse,
	fromCache bool,
) {
	ctx.JSON(
		http.StatusOK,
		struct {
			models.UserResponse
			FromCache bool `json:"fromCache"`
		}{
			UserResponse: userResponse,
			FromCache:    fromCache,
		},
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
		gin.H{"user": userResponse},
	)
}
