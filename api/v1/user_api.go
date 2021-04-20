package v1

import (
	helpers "firmanjs/example-gin/helpers"
	model "firmanjs/example-gin/repository/models"
	req "firmanjs/example-gin/repository/request/v1"
	service "firmanjs/example-gin/services/v1"
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetAllUsers ..
func GetAllUsers(c *gin.Context) {

	response := service.GetAll()

	helpers.Respond(c.Writer, response)
}

//CreateUser ....
func CreateUser(c *gin.Context) {
	// idUser := c.MustGet("credUser").(string)
	idUser := "System"
	var user *req.CreateUserReq
	c.BindJSON(&user)

	response := service.Create(idUser, user)

	helpers.Respond(c.Writer, response)
}

//UpdateUser ..
func UpdateUser(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var user *req.EditUserReq
	c.BindJSON(&user)

	response := service.Update(idUser, user)

	helpers.Respond(c.Writer, response)
}

//GetUser ..
func GetUser(c *gin.Context) {
	//userID := c.Query("id")
	idUser := c.MustGet("credUser").(string)
	response := service.GetByID(idUser)
	fmt.Println(idUser)
	helpers.Respond(c.Writer, response)
}

//DeleteUser ..
func DeleteUser(c *gin.Context) {
	userID := c.Query("id")
	idUser := c.MustGet("credUser").(string)
	response := service.DeleteByID(idUser, userID)

	helpers.Respond(c.Writer, response)
}

//LoginUser ..
func LoginUser(c *gin.Context) {
	var user *model.Users
	c.BindJSON(&user)

	response := service.Login(user)

	helpers.Respond(c.Writer, response)
}

//ChangePassword ...
func ChangePassword(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var reqModel *req.ChangePassReqModel
	c.BindJSON(&reqModel)

	response := service.ChangePass(idUser, reqModel)

	helpers.Respond(c.Writer, response)

}
