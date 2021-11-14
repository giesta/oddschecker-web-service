package main

import (
	"oddschecker/oddschecker-web-service/controllers"
	"oddschecker/oddschecker-web-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.POST("/bets", controllers.CreateBet)
	r.GET("/bets", controllers.GetBets)
	r.GET("/bets/:id", controllers.GetBet)
	r.PATCH("/bets/:id", controllers.UpdateBet)
	r.DELETE("/bets/:id", controllers.DeleteBet)

	r.POST("/odds", controllers.CreateOdd)
	r.GET("/odds/:betId", controllers.GetOdds)

	r.Run()
}
