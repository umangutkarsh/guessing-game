// router.go
package bootstrap

import (
	"path"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	api := "api"
	version := "v1"

    r := gin.Default()

	r.RedirectTrailingSlash = false

	// Enable CORS with custom configuration.
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://guessing-game-ffjoqgttq-umangutkarshs-projects.vercel.app"}, // Frontend URL from environment variables
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

    // Define routes, for example:
    // r.GET("/ping", func(c *gin.Context) {
    //     c.JSON(200, gin.H{"message": "pong"})
    // })

	routerGroup := r.Group(path.Join(api, version))
	

	SetupRoutes(routerGroup)

    return r
}
