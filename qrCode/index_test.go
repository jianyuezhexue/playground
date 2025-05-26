package qrCode

import (
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestCreateQrCodeWithLogo(t *testing.T) {
	// 创建一个二维码图片，并添加logo
	err := CreateQrCodeWithLogo("https://www.pgyer.com/0VH3LLBi?ZHCSXYZH=bZ16xTYW6vmucQhqZMrYIM0/aIgwA2UNQftpZL/h/t8XVhrdZrYhJKAhRIyJl1CU7ApDrsmpQynNxHmJp7wfBgWvVsfE4I0WzWXUONBv9SBN7PPMneDWyPbP7ukRIKQa9Kr0eT6eNI/3yfPk1tcWN5qDSbLG0HnSUghnc8rUev27M8/kBBkHQb+OGsxF62U63qZtRCoBUUHqyX5RgWg890pQsSMdHXj80TzXZoS0Uls=", "C:\\Users\\jason_wang2\\Desktop\\guoneng_logo.png",
		"./qrCode4.png", qrcode.Medium, 256)
	if err != nil {
		t.Errorf("Failed to create QR code with logo: %v", err)
	}
}
