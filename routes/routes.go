package routes

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	handlers "github.com/ruiblaese/projeto-codenation-banco-uati/handlers"
	middleware "github.com/ruiblaese/projeto-codenation-banco-uati/middleware"
)

//StartRouter inicia servidor e estabelece rotas
func StartRouter() {
	router := gin.Default()

	authMiddleware := middleware.GetAuthMiddleware()

	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", handlers.HelloHandler)
	}

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	//exemplos
	auth.GET("/jokes", handlers.JokeHandler)
	auth.POST("/jokes/like/:jokeID", handlers.LikeJoke)

	//autenticar usuario
	auth.POST("/signin/", handlers.LikeJoke)

	//cadastrar usuario
	auth.POST("/signup/", handlers.LikeJoke)

	//inicia servidor
	router.Run(":3000")

}
