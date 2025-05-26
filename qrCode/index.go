package qrCode

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/draw"
)

// GetQRCodeIO 返回图片字节 content-二维码内容   level-容错级别(越高越好),Low,Medium,High,Highest   size-像素单位
func GetQRCodeIO(content string, level qrcode.RecoveryLevel, size int) string {
	var png []byte
	png, err := qrcode.Encode(content, level, size)
	if err != nil {
		return ""
	}
	res := base64.StdEncoding.EncodeToString(png)
	fmt.Println(res)
	return res
}

// GetQRCodeFile content-二维码内容   level-容错级别,Low,Medium,High,Highest   size-像素单位  outPath-输出路径
func GetQRCodeFile(content, outPath string, level qrcode.RecoveryLevel, size int) interface{} {
	err := qrcode.WriteFile(content, level, size, outPath)
	if err != nil {
		return err.Error()
	}
	return nil
}

// GetQRCodeCustom content-二维码内容   level-容错级别,Low,Medium,High,Highest   size-像素单位  outPath-输出路径   bColor-前景颜色   gColor-背景颜色
func GetQRCodeCustom(content, outPath string, level qrcode.RecoveryLevel, size int, bColor, gColor color.Color) interface{} {
	err := qrcode.WriteColorFile(content, level, size, bColor, gColor, outPath)
	if err != nil {
		return err.Error()
	}
	return nil
}

// CreateQrCodeWithLogo 带logo的二维码图片生成 content-二维码内容   level-容错级别,Low,Medium,High,Highest   size-像素单位  outPath-输出路径  logoPath-logo文件路径
func CreateQrCodeWithLogo(content, logoPath, outPath string, level qrcode.RecoveryLevel, size int) error {
	code, err := qrcode.New(content, level)
	if err != nil {
		return err
	}
	//设置文件大小并创建画板
	qrcodeImg := code.Image(size)
	outImg := image.NewRGBA(qrcodeImg.Bounds())

	//读取logo文件
	logoFile, err := os.Open(logoPath)
	if err != nil {
		panic(err)
	}
	logoImg, _, err := image.Decode(logoFile)
	logoImg = resize.Resize(uint(size/6), uint(size/6), logoImg, resize.Lanczos3)

	//logo和二维码拼接
	draw.Draw(outImg, outImg.Bounds(), qrcodeImg, image.Pt(0, 0), draw.Over)
	offset := image.Pt((outImg.Bounds().Max.X-logoImg.Bounds().Max.X)/2, (outImg.Bounds().Max.Y-logoImg.Bounds().Max.Y)/2)
	draw.Draw(outImg, outImg.Bounds().Add(offset), logoImg, image.Pt(0, 0), draw.Over)

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	png.Encode(f, outImg)
	return nil
}

// CreateQrCodeCustomWithLogo 带logo的二维码图片生成 content-二维码内容   level-容错级别,Low,Medium,High,Highest   size-像素单位  outPath-输出路径  logoPath-logo文件路径
func CreateQrCodeCustomWithLogo(content, logoPath, outPath string, level qrcode.RecoveryLevel, size int) interface{} {
	code, err := qrcode.New(content, level)
	if err != nil {
		return err.Error()
	}
	qrcodeImg := code.Image(size)
	outImg := image.NewRGBA(qrcodeImg.Bounds())

	logoFile, err := os.Open(logoPath)
	if err != nil {
		panic(err)
	}
	logoImg, _, err := image.Decode(logoFile)
	logoImg = resize.Resize(uint(size/6), uint(size/6), logoImg, resize.Lanczos3)

	//添加方形画板
	circleImg := code.Image(size/6 + 5)
	outCircleImg := image.NewRGBA(circleImg.Bounds())
	//logo切为圆形
	draw.DrawMask(outCircleImg, outCircleImg.Bounds(), logoImg, image.ZP, &Circle{image.Pt(size/12, size/12), size / 12}, image.ZP, draw.Over)

	//logo和二维码拼接
	draw.Draw(outImg, outImg.Bounds(), qrcodeImg, image.Pt(0, 0), draw.Over)
	offset := image.Pt((outImg.Bounds().Max.X-outCircleImg.Bounds().Max.X)/2, (outImg.Bounds().Max.Y-logoImg.Bounds().Max.Y)/2)
	draw.Draw(outImg, outImg.Bounds().Add(offset), outCircleImg, image.Pt(0, 0), draw.Over)

	//再次添加画板
	backImg := code.Image(size - size/10)
	outBackImg := image.NewRGBA(backImg.Bounds())
	draw.DrawMask(outBackImg, outBackImg.Bounds(), outImg, image.ZP, &Rectangle{image.Pt((size+(size/10))/2, (size+(size/10))/2), size/2 - size/20, size/2 - size/20}, image.ZP, draw.Over)

	f, err := os.Create(outPath)
	if err != nil {
		return err.Error()
	}
	png.Encode(f, outBackImg)
	return nil
}
