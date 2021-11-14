package main

import (
	"github.com/giesta/oddschecker-web-service/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Run()
}
