package controllers

import (
	"net/http"

	"oddschecker/oddschecker-web-service/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}
type UpdateUserRequest struct {
	Name string `json:"name"`
}

// POST /users
func CreateUser(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user format"})
		return
	}

	id := uuid.New()

	user := models.User{
		ID:   id.String(),
		Name: request.Name,
	}
	result := models.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user format"})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// GET /users
// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Get an user
func GetUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
// Update an user
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var request UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := models.DB.Model(&user).Updates(request)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user format"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete an user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	result := models.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't delete"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
