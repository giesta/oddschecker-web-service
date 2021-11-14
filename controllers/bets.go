package controllers

import (
	"net/http"

	"oddschecker/oddschecker-web-service/models"

	"github.com/gin-gonic/gin"
)

type CreateBetRequest struct {
	Name string `json:"name" binding:"required"`
}
type UpdateBetRequest struct {
	Name string `json:"name"`
}

// POST /bets
func CreateBet(c *gin.Context) {
	var request CreateBetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bet format"})
		return
	}

	bet := models.Bet{
		Name: request.Name,
	}
	result := models.DB.Create(&bet)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bet format"})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// GET /bets
// Get all bets
func GetBets(c *gin.Context) {
	var bets []models.Bet
	models.DB.Find(&bets)

	c.JSON(http.StatusOK, gin.H{"data": bets})
}

// GET /bets/:id
// Get a bet
func GetBet(c *gin.Context) {
	var bet models.Bet

	if err := models.DB.Where("id = ?", c.Param("id")).First(&bet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bet not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bet})
}

// PATCH /bets/:id
// Update a bet
func UpdateBet(c *gin.Context) {
	var bet models.Bet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&bet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bet not found!"})
		return
	}

	var request UpdateBetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := models.DB.Model(&bet).Updates(request)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bet format"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bet})
}

// DELETE /bet/:id
// Delete a bet
func DeleteBet(c *gin.Context) {
	var bet models.Bet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&bet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bet not found!"})
		return
	}

	result := models.DB.Delete(&bet)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't delete"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
