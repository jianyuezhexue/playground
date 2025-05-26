package main

import (
	"fmt"
	"time"

	"github.com/jianyuezhexue/playground/qrCode"
	"github.com/skip2/go-qrcode"
)

func main() {
	lastMaxCreateTime := time.Now().Format("2006-01-02 00:00:00")
	fmt.Println(lastMaxCreateTime)

	//这里我用的gin框架，框架不重要，重要的是使用
	// qrCode.GetQRCodeIO("https://www.baidu.com", qrcode.Medium, 256)
	// qrCode.GetQRCodeFile("https://www.baidu.com", "./qrCode1.png", qrcode.Medium, 256)
	// qrCode.GetQRCodeCustom("https://www.baidu.com", "config/qrCode2.png", qrcode.Medium, 256,
	// 	color.RGBA{R: 50, G: 50, B: 50, A: 50}, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	qrCode.CreateQrCodeWithLogo("https://www.pgyer.com/0VH3LLBi?ZHCSXYZH=bZ16xTYW6vmucQhqZMrYIM0/aIgwA2UNQftpZL/h/t8XVhrdZrYhJKAhRIyJl1CU7ApDrsmpQynNxHmJp7wfBgWvVsfE4I0WzWXUONBv9SBN7PPMneDWyPbP7ukRIKQa9Kr0eT6eNI/3yfPk1tcWN5qDSbLG0HnSUghnc8rUev27M8/kBBkHQb+OGsxF62U63qZtRCoBUUHqyX5RgWg890pQsSMdHXj80TzXZoS0Uls=", "C:\\Users\\jason_wang2\\Desktop\\guoneng_logo.png",
		"./qrCode3.png", qrcode.High, 256)
	// qrCode.CreateQrCodeCustomWithLogo("https://www.baidu.com", "E:\\图片\\头像\\2a085da2a850392d0d2b6d840d4dc4e5.jpeg",
	// 	"config/qrCode4.png", qrcode.Medium, 256)
}
