package handlers

import (
	"fmt"
	"net/http"

	"guessing-game/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ChallengeService *services.ChallengeService

func CreateChallengeHandler(c *gin.Context) {
	var req struct {
		UserID string `json:"userId" binding:"required"`
	}
	fmt.Println("creating challenge...")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("binded")
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}
	fmt.Println("user id invalid...")
	challenge, err := ChallengeService.CreateChallenge(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("created challenge...")
	inviteImageURL := "https://cdn5.vectorstock.com/i/1000x1000/89/89/invite-friends-rubber-stamp-vector-12798989.jpg"

	c.JSON(http.StatusOK, gin.H{
		"challengeId": challenge.Token,
		"inviteLink":  challenge.InviteLink,
		"inviteImage": inviteImageURL,
	})
}

func GetChallengeHandler(c *gin.Context) {
	token := c.Param("challengeId")
	challenge, err := ChallengeService.GetChallenge(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	inviter, err := AuthService.GetProfile(challenge.InviterUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "inviter not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"challengeId":    challenge.Token,
		"inviteLink":     challenge.InviteLink,
		"inviterUserId":  challenge.InviterUserID.String(),
		"inviterUsername": inviter.Username,
		"inviterScore":   inviter.Score,
	})
}
