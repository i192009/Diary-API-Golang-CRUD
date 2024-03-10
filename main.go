package main

import (
	"diaryApi/controller"
	Database "diaryApi/database"
	"diaryApi/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()
	loadDatabase() 
	serveApplication()
}

func loadDatabase() {
	Database.Connect()
	if err := Database.Database.AutoMigrate(&model.User{}); err != nil {
		log.Println(err)
	}
	if err := Database.Database.AutoMigrate(&model.Entry{}); err != nil {
		log.Println(err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register" , controller.Register) 
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.GET("/users" , controller.GetUsers)

	router.Run(":8000")
	fmt.Println("Server is running on port 8000")
}

