package main

import (
	"context"
	//"fmt"
	//"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "ExpenceTracker/docs"
	"ExpenceTracker/internal/config"
	"ExpenceTracker/internal/controller"
	"ExpenceTracker/internal/db"
	"ExpenceTracker/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	// Загрузка конфигов и логгера
	config.LoadConfig()
	logger.InitLogger()

	// .env
	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("⚠️  Не удалось загрузить .env файл")
	}

	// Подключение к БД
	if err := db.ConnectDB(); err != nil {
		logger.Log.Fatal("❌ Ошибка подключения к базе данных: ", err)
	}
	defer db.CloseDB()

	// Миграции
	if err := db.InitMigrations(); err != nil {
		logger.Log.Fatal("❌ Ошибка миграции: ", err)
	}

	// Роутинг
	router := gin.Default()
	controller.RegisterRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// HTTP-сервер
	srv := &http.Server{
		Addr:    ":8181",
		Handler: router,
	}

	logger.Log.Info("✅ Инфо-сообщение")
	logger.Log.Warn("⚠️ Это предупреждение")
	logger.Log.Error("❌ Ошибка какая-то")
	logger.Log.Debug("🛠 Это отладочная информация")

	// Асинхронный запуск
	go func() {
		logger.Log.Info("🚀 Сервер запущен на http://localhost:8181")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("❌ Ошибка сервера: ", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("⏳ Завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("❌ Ошибка при завершении работы сервера: ", err)
	}

	logger.Log.Info("✅ Сервер успешно остановлен")

}
