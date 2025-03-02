// internal/bootstrap/routes.go
package bootstrap

import (
	"guessing-game/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	// Auth routes
	auth := r.Group("/auth")

	auth.POST("/register", handlers.RegisterUserHandler)
	auth.GET("/profile/:userId", handlers.GetProfileHandler)
	auth.GET("/profile/username/:username", handlers.GetProfileByUsernameHandler)

	// Game routes
	game := r.Group("/game")

	game.GET("/destination", handlers.GetRandomDestinationHandler)
	game.POST("/guess", handlers.GuessHandler)
	game.GET("/score/:userId", handlers.GetScoreHandler)

	// Challenge routes
	challenge := r.Group("/challenge")

	challenge.POST("", handlers.CreateChallengeHandler)
	challenge.GET("/:challengeId", handlers.GetChallengeHandler)
}

