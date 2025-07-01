package main

import (
	_ "ExpenceTracker/docs"
	"ExpenceTracker/internal/controller"
	"ExpenceTracker/internal/db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Expense Tracker API
// @version 1.0
// @description API Server for Expense Tracker Application
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8181
// @BasePath /

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	if err := db.ConnectDB(); err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer db.CloseDB()

	if err := db.InitMigrations(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	router := gin.Default()

	controller.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    ":8181",
		Handler: router,
	}

	go func() {
		fmt.Println(" Server running at http://localhost:8181")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown failed: %v", err)
	}

	fmt.Println("✅ Server stopped properly")
}
