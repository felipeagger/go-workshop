package main

import (
	"fmt"
	"os"

	"go-workshop/src"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	health := engine.Group("/")

	userRoute := engine.Group("/user")

	//User
	userRoute.GET("/", src.GetUsers)
	userRoute.GET("/:id", src.GetUser)
	userRoute.POST("/", src.PostUser)
	userRoute.PUT("/:id", src.PutUser)
	userRoute.DELETE("/:id", src.DeleteUser)

	src.AutoMigration()

	health.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Go healthy!",
		})
	})

	engine.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
