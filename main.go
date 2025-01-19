package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haniifac/be-user-service/controller"
)

func main() {
	router := gin.Default()

	userGroup := router.Group("/users")
	userGroup.GET("/", controller.GetUsers)
	userGroup.POST("/", controller.CreateUser)
	userGroup.PUT("/:id", controller.UpdateUser)
	userGroup.DELETE("/:id", controller.DeleteUser)

	router.Run()
}
