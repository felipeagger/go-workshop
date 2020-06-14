package main

import (
	"fmt"
	"os"

	"github.com/felipeagger/go-workshop/src"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	health := engine.Group("/")

	userRoute := engine.Group("/user")
	debtRoute := engine.Group("/debt")

	src.AutoMigration()

	//User
	userRoute.GET("/", src.GetUsers)
	userRoute.GET("/:id", src.GetUser)
	userRoute.POST("/", src.PostUser)
	userRoute.PUT("/:id", src.PutUser)
	userRoute.DELETE("/:id", src.DeleteUser)

	userRoute.GET("/:id/debts", src.GetDebtsByUser)

	//Debt
	debtRoute.GET("/", src.GetDebts)
	debtRoute.GET("/:id", src.GetDebt)
	debtRoute.POST("/", src.PostDebt)
	debtRoute.PUT("/:id", src.PutDebt)
	debtRoute.DELETE("/:id", src.DeleteDebt)

	health.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go healthy!",
		})
	})

	engine.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))

}
