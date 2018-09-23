package main

import (
	"github.com/gin-gonic/gin"

	"gin_api/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/users", controllers.Index)
	router.GET("/users/:id", controllers.Show)
	router.POST("/users", controllers.Create)
	router.DELETE("/users/:id", controllers.Delete)
	router.Run(":8080")
}
