package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jobayer12/go-crud/controllers"
	"github.com/jobayer12/go-crud/models"
	"log"
)

func main() {
	server := gin.Default()
	models.ConnectDatabase()

	router := server.Group("/api/v1")

	router.GET("user", controllers.Users)
	router.POST("user", controllers.CreateUser)
	router.PUT("user/:id", controllers.UpdateUserById)
	router.DELETE("user/:id", controllers.DeleteUserById)
	log.Fatal(server.Run())
}
