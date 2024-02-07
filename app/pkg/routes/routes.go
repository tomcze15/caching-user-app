package routes

import (
	"github.com/gin-gonic/gin"

	"caching-user-app/pkg/controllers"
)

const (
	baseApi = "/users"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET(baseApi+"/:id", controllers.GetUser)
	router.GET(baseApi, controllers.GetUsers)
	router.POST(baseApi, controllers.CreateUser)

	return router
}
