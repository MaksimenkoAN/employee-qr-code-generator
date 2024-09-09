package qr

import (
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(data string) ([]byte, error) {
	return qrcode.Encode(data, qrcode.Medium, 256)
}
