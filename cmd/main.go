package main

import (
	"log"

	v1 "employee-qr-code-generator/api/v1"
	"employee-qr-code-generator/internal/config"
	"employee-qr-code-generator/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	err := config.LoadConfig("../configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Инициализация БД
	err = database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	r := gin.Default()

	// Регистрация маршрутов API
	v1.RegisterRoutes(r)

	// Определяем хост и порт
	addr := ":8080"

	// Запуск сервера
	if err := r.Run(addr); err != nil {
		panic(err)
	}
}
