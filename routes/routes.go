package routes

import (
	"controllers"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	// Our API will consit of just two routes
	// /jokes - which will retrieve a list of jokes a user can see
	// /jokes/like/:jokeID - which will capture likes sent to a particular joke
	api.GET("/jokes", controllers.JokeHandler)
	api.POST("/jokes/like/:jokeID", controllers.LikeJoke)

	// Start and run the server
	router.Run(":3000")

}
