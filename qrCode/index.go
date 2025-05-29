package qrCode

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"

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

// CreateQrCodeWithLogoBase64 带logo的二维码图片生成Base64字符串 content-二维码内容   level-容错级别,Low,Medium,High,Highest   size-像素单位   logoPath-logo文件路径
func CreateQrCodeWithLogoBase64(content, logoPath string, level qrcode.RecoveryLevel, size int) (string, error) {
	// 1. 生成二维码
	code, err := qrcode.New(content, level)
	if err != nil {
		return "", err
	}

	// 2. 创建二维码图像
	qrcodeImg := code.Image(size)
	outImg := image.NewRGBA(qrcodeImg.Bounds())

	// 3. 读取Logo文件
	logoFile, err := ReadImage(logoPath) // 仍然接收文件路径
	if err != nil {
		return "", err
	}

	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		return "", err
	}
	logoImg = resize.Resize(uint(size/5), uint(size/5), logoImg, resize.Lanczos3)

	// 4. 合成二维码和Logo
	draw.Draw(outImg, outImg.Bounds(), qrcodeImg, image.Pt(0, 0), draw.Over)
	offset := image.Pt((outImg.Bounds().Max.X-logoImg.Bounds().Max.X)/2, (outImg.Bounds().Max.Y-logoImg.Bounds().Max.Y)/2)
	draw.Draw(outImg, outImg.Bounds().Add(offset), logoImg, image.Pt(0, 0), draw.Over)

	// 5. 转换为Base64
	var buf bytes.Buffer
	if err := png.Encode(&buf, outImg); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// CreateQrCodeWithLogo 带logo的二维码图片生成 content-二维码内容   level-容错级别,Low,Medium,High,Highest   size-像素单位  outPath-输出路径  logoPath-logo文件路径
func CreateQrCodeWithLogo(content, logoPath, outPath string, level qrcode.RecoveryLevel, size int) error {
	code, err := qrcode.New(content, level)
	if err != nil {
		return err
	}

	// 设置文件大小并创建画板
	qrcodeImg := code.Image(size)
	outImg := image.NewRGBA(qrcodeImg.Bounds())

	// 读取Logo文件
	logoFile, err := ReadImage(logoPath)
	if err != nil {
		return err
	}
	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		return err
	}
	logoImg = resize.Resize(uint(size/6), uint(size/6), logoImg, resize.Lanczos3)

	// Logo和二维码拼接
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

// ReadImage 返回 io.Reader，兼容本地文件和远程 URL
func ReadImage(pathOrURL string) (io.Reader, error) {
	if strings.HasPrefix(pathOrURL, "http://") || strings.HasPrefix(pathOrURL, "https://") {
		// 远程 URL
		return readFromURL(pathOrURL)
	} else {
		// 本地文件
		file, err := os.Open(pathOrURL)
		if err != nil {
			return nil, fmt.Errorf("打开文件失败: %v", err)
		}
		return file, nil // 返回 *os.File，调用方需自行关闭
	}
}

// 从远程 URL 读取，返回 io.ReadCloser（实现了 io.Reader）
func readFromURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP 请求失败: %v", err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // 必须关闭，否则资源泄漏
		return nil, fmt.Errorf("HTTP 请求失败，状态码: %d", resp.StatusCode)
	}

	return resp.Body, nil
}
