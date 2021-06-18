package api

import (
	"github.com/dscamargo/cobraapp/api/controllers"
	"github.com/dscamargo/cobraapp/app/handlers"
	"github.com/gin-gonic/gin"
)

func configRoutes(api *gin.RouterGroup) {
	userService := handlers.NewUserHandler()
	userController := controllers.NewController(userService)

	api.POST("/users", userController.CreateUser)
	api.GET("/users", userController.ListUsers)
}

func Start(port string) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("api")
	configRoutes(api)
	r.Run(":" + port)
}
