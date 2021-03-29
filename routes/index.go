package routers

import (
    "net/http"
	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {

	routes := gin.Default()

    routes.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
    })

	return routes
}