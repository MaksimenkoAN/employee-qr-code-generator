package handler

import (
	"employee-qr-code-generator/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var idViewAllInfo int = 2280

// GenerateContactQRCode обрабатывает запросы на генерацию QR-кодов для добавления контактов
func GenerateContactQRCode(c *gin.Context) {
	//employeeID := c.Query("employeeID")
	//var addPrivileges bool = false
	var username string = c.Query("username")
	userIDs := database.GetUserId(username)
	for _, id := range userIDs {
		if id == idViewAllInfo {
			//addPrivileges = true
		}
	}
	//addPrivileges = false

	c.IndentedJSON(http.StatusOK, userIDs)
	/*
		// Пример получения информации о сотруднике из базы данных
		// Здесь должно быть вызов к функции из database, чтобы получить детали сотрудника
		name := "John Doe" // Пример данных
		phone := "+1234567890"
		email := "john.doe@example.com"

		contactInfo := fmt.Sprintf(
			"BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nEMAIL:%s\nEND:VCARD",
			name, phone, email,
		)

		// Генерация QR-кода
		qr, err := qrcode.Encode(contactInfo, qrcode.Medium, 256)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
			return
		}

		// Установка заголовков и отправка QR-кода как изображения
		c.Header("Content-Type", "image/png")
		c.Writer.Write(qr)*/
}
