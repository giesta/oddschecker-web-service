package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"oddschecker/oddschecker-web-service/models"

	"github.com/gin-gonic/gin"
)

type CreateOddRequest struct {
	BetId  *uint  `json:"betId" binding:"required"`
	UserId string `json:"userId" binding:"required"`
	Odds   string `json:"odds" binding:"required"`
}

// POST /odds
// Offer odds for a bet
func CreateOdd(c *gin.Context) {
	// Validate request
	var request CreateOddRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format of Odds"})
		return
	}

	rxOdds := regexp.MustCompile(`^([1-9][0-9]*\/[0-9]+|SP)$`)

	if !rxOdds.MatchString(request.Odds) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format of Odds"})
		return
	}

	// Create odd
	odd := models.Odd{
		BetId:  *request.BetId,
		UserId: request.UserId,
		Odds:   request.Odds,
	}
	result := models.DB.Create(&odd)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format of Odds"})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// GET /odds/:betId
// Find Odds by Bet ID
func GetOdds(c *gin.Context) {
	var odds []models.Odd

	betId := c.Param("betId")

	if len(betId) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Bet ID supplied"})
		return
	}

	_, err := strconv.Atoi(betId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Bet ID supplied"})
		return
	}

	var bet models.Bet

	if err := models.DB.Where("id = ?", betId).First(&bet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bet not found for given ID"})
		return
	}

	if err := models.DB.Where("bet_id = ?", bet.ID).Find(&odds).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bet not found for given ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": odds})
}
