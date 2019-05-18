module projeto-codenation-banco-uati

go 1.12

require (
	controllers v0.0.0
	db v0.0.0
	github.com/gin-gonic/contrib v0.0.0-20190510065052-87e961e51ccc
	github.com/gin-gonic/gin v1.4.0
	models v0.0.0
	routes v0.0.0

)

replace (
	controllers => ./controllers
	db => ./db
	models => ./models
	routes => ./routes
	services => ./services

)
