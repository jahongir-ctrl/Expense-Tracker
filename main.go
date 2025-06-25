package main

import (
	"ExpenceTracker/internal/controller"
	"ExpenceTracker/internal/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Expense Tracker API is running"})
	//})

	r.Run(":8181") // Run on port 8181
	fmt.Println("Expense Tracker API is running on http://localhost:8181")
}
