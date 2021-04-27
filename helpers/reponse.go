package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseData structure
type ResponseData struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

//Message return map data
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond return basic response structure
func Respond(c *gin.Context, code int, status string, message string, data interface{}) {
	c.JSON(code, gin.H{
		"data":    data,
		"status":  status,
		"message": message,
	})
}

func RespondOk(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  "Success",
		"message": message,
	})
}

func RespondError(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    nil,
		"status":  "Error",
		"message": message,
	})
}

func RespondNotFOund(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusNotFound, gin.H{
		"data":    nil,
		"status":  "Not found",
		"message": "Data not found",
	})
}

func RespondBad(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"data":    nil,
		"status":  "Error",
		"message": message,
	})
}

func RespondCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"status":  "Created",
		"message": "Data successfully created",
	})
}
