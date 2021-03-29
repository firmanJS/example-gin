package main

import (
	"fmt"
	"os"
	"firmanjs/example-gin/routes"
	"github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	r := routers.SetupRouter()

	port := "8080"
    gin.SetMode("debug")
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8000"
	}

	type Job interface {
		Run()
	}

	r.Run(":" + port)

}