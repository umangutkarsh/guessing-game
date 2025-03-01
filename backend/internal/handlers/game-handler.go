package handlers

import (
	"net/http"

	"guessing-game/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Exported service instance; will be set in main.
var GameService *services.GameService

func GetRandomDestinationHandler(c *gin.Context) {
	dest, err := GameService.GetRandomDestination()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dest)
}

func GuessHandler(c *gin.Context) {
	var req struct {
		UserID         string `json:"userId"`
		QuestionToken  string `json:"questionToken"`
		SelectedCityID string `json:"selectedCityId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}
	selectedCityID, err := uuid.Parse(req.SelectedCityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid selectedCityId"})
		return
	}

	result, err := GameService.ProcessGuess(userID, req.QuestionToken, selectedCityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetScoreHandler(c *gin.Context) {
	userIdStr := c.Param("userId")
	userID, err := uuid.Parse(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}
	score, err := GameService.GetScore(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userId": userID.String(), "score": score})
}
