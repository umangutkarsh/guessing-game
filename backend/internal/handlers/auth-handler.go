package handlers

import (
	"net/http"

	"guessing-game/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Exported service instance; will be set in main.
var AuthService *services.AuthService

func RegisterUserHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := AuthService.RegisterUser(req.Username)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId":   user.ID.String(),
		"username": user.Username,
		"score":    user.Score,
	})
}

func GetProfileHandler(c *gin.Context) {
	userIdStr := c.Param("userId")
	userID, err := uuid.Parse(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}
	user, err := AuthService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId":   user.ID.String(),
		"username": user.Username,
		"score":    user.Score,
	})
}

func GetProfileByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	user, err := AuthService.GetProfileByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId":   user.ID.String(),
		"username": user.Username,
		"score":    user.Score,
	})
}