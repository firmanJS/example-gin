package routers

import (
	api "firmanjs/example-gin/api/v1"

	"github.com/gin-gonic/gin"
)

func routeUser(routes *gin.Engine) {

	user := routes.Group("/api/v1/user")

	user.POST("/", api.CreateUsers)
	user.GET("/", api.GetAllUsers)
	user.GET("/:id", api.GetUsers)
	user.DELETE("/:id", api.DeleteUsers)
}
