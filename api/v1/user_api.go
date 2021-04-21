package v1

import (
	helpers "firmanjs/example-gin/helpers"
	req "firmanjs/example-gin/repository/request/v1"
	service "firmanjs/example-gin/services/v1"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	code, status, message, data := service.GetAll()

	helpers.Respond(c, code, status, message, data)
}

func CreateUsers(c *gin.Context) {
	idUser := "System"
	var user *req.CreateUserReq
	c.BindJSON(&user)

	code, status, message, data := service.Create(idUser, user)

	helpers.Respond(c, code, status, message, data)
}

func UpdateUsers(c *gin.Context) {
	// idUser := c.MustGet("credUser").(string)
	// var user *req.EditUserReq
	// c.BindJSON(&user)

	// response := service.Update(idUser, user)

	// helpers.Respond(c.Writer, response)
}

func GetUsers(c *gin.Context) {
	id := c.Param("id")
	code, status, message, data := service.GetByParam(id)
	helpers.Respond(c, code, status, message, data)

	fmt.Println(id)
}

func DeleteUsers(c *gin.Context) {
	id := c.Param("id")

	code, status, message, data := service.DeleteByID(id)

	helpers.Respond(c, code, status, message, data)
}

