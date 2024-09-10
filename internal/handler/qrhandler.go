package handler

import (
	"employee-qr-code-generator/internal/database"
	"employee-qr-code-generator/pkg/qr"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var idViewAllInfo int = 2280

// GenerateContactQRCode обрабатывает запросы на генерацию QR-кодов для добавления контактов
func GenerateContactQRCode(c *gin.Context) {
	var addPrivileges bool = false
	employeeID := c.Query("employeeID")
	var username string = c.Query("username")
	userIDs := database.GetUserId(username)

	for _, id := range userIDs {
		if id == idViewAllInfo {
			addPrivileges = true
		}
	}
	employeeInfo := database.GetInfoEmployee(employeeID, addPrivileges)
	name := employeeInfo.Name
	email := employeeInfo.Email
	workAddress := employeeInfo.Address
	mobilePhone := employeeInfo.MobilePhone
	str := strings.ReplaceAll(employeeInfo.WorkPhone, " ext. ", ",,")
	str = strings.ReplaceAll(str, " ", "")
	workPhone := str

	contactInfo := fmt.Sprintf(
		"BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL;TYPE=WORK:%s\nTEL;TYPE=CELL:%s\nADR;TYPE=WORK:%s\nEMAIL:%s\nEND:VCARD",
		name, workPhone, mobilePhone, workAddress, email,
	)

	// Генерация QR-кода
	qr, err := qr.GenerateQRCode(contactInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	// Установка заголовков и отправка QR-кода как изображения
	c.Header("Content-Type", "image/png")
	c.Writer.Write(qr)
}
