package officetopdf

import (
	"log"
	"os/exec"
)

// 转换文件为PDF格式
// libreoffice官网地址： https://zh-cn.libreoffice.org/
// 配置环境变量：在系统环境变量中添加路径，例如：C:\Program Files\LibreOffice\program
// 安装好后，在命令行输入：soffice --version 查看是否安装成功
// 使用方法： soffice --headless --convert-to pdf input.docx --outdir output_folder/
func ConvertToPDF(inputFile, outputFile string) {
	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", inputFile, "--outdir", outputFile)
	err := cmd.Run()
	if err != nil {
		log.Fatal("转换失败：", err)
	}
	log.Println("转换成功！")
}
