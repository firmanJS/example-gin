package routers

import (
	db "firmanjs/example-gin/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {
	//connecting to db
	db.Connect()

	//routing endpoint

	routes := gin.Default()

	routes.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	return routes
}
