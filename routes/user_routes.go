package routers

import (
	api "firmanjs/example-gin/api/v1"

	"github.com/gin-gonic/gin"
)

func routeUser(routes *gin.Engine) {

	user := routes.Group("/api/v1/user")

	user.POST("/login", api.LoginUser)
	user.POST("/", api.CreateUser)
	user.GET("/all", api.GetAllUsers)
}
