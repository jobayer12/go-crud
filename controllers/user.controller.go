package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jobayer12/go-crud/models"
	"net/http"
)

func Users(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func CreateUser(c *gin.Context) {
	// Validate input
	var input models.CreateUserRequestPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetByUserId(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id!"})
		return
	}

	var input models.UpdateUserRequestPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id!"})
		return
	}

	models.DB.Delete(&user, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
