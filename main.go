package main

import (
	"latihangin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"sukses": 1, "pesan": "Halo Dunia dari Ngkong Sayid!!"})
	})

	router.POST("/users/register", controllers.RegisterUser())
	router.GET("/users", controllers.ListUser())
	router.GET("/users/:id", controllers.DetailUser())
	router.DELETE("/users/:id", controllers.HapusUser())
	router.PUT("/users/:id", controllers.UbahUser())
	router.POST("/users/login", controllers.LoginUser())

	router.Run() // ":8080" untuk ganti port
}
