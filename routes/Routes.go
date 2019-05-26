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

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.POST("/api/v1/signin", authMiddleware.LoginHandler)
	router.POST("/api/v1/signup", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1 := router.Group("/api/v1")
	// Refresh time can be longer than token timeout
	v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/hello", handlers.HelloHandler)

		//rotas para usuarios
		user := v1.Group("/user")
		{
			user.GET("", handlers.GetUsers)
			user.GET(":id", handlers.GetUser)
			user.PUT(":id", handlers.PutUser)
			user.POST("", handlers.NewUser)
			user.DELETE(":id", handlers.DeleteUser)
		}

		//rotas para clientes
		customer := v1.Group("/customer")
		{
			customer.GET("", handlers.GetCustomers)
			customer.GET(":id", handlers.GetCustomer)
			customer.PUT(":id", handlers.PutCustomer)
			customer.POST("", handlers.NewCustomer)
			customer.DELETE(":id", handlers.DeleteCustomer)
			customer.POST("upload", handlers.UploadCustomersWithCSV)
		}
	}

	//inicia servidor
	router.Run(":3000")

}
