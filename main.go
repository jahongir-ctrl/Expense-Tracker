package main

import (
	_ "ExpenceTracker/docs"
	"ExpenceTracker/internal/controller"
	"ExpenceTracker/internal/db"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MindsWallet API
// @version 1.0
// @description API Server for MindsWallet Application
// @securityDefinitions.apikey ApiKeyAuth
// @host localhost:8181
// @BasePath /
// @in header
// @name Authorization

func main() {

	err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error to connect with DB :", err)
	}
	defer db.CloseDB()

	err = db.InitMigrations()
	if err != nil {
		fmt.Println("Error to init migration :", err)
	}

	r := gin.Default()
	controller.RegisterRoutes(r)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Expense Tracker API is running"})
	//})

	r.Run(":8181") // Run on port 8181
	fmt.Println("Expense Tracker API is running on http://localhost:8181")
}
