package v1

import (
	"employee-qr-code-generator/internal/handler" // Логика обработки QR-кодов

	"github.com/gin-gonic/gin"
)

// RegisterRoutes регистрирует маршруты API
func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/generate-contact-qr/", handler.GenerateContactQRCode)
	}
}
