package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main5() {
	router := gin.Default()
	router.GET("/time", func(c *gin.Context) {
		m := map[string]string{
			"time": time.Now().Format("2006-01-02"),
		}
		c.JSON(http.StatusOK, m)
	})
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}

//func main() {
//	router := gin.Default()
//
//	// 设置上传文件的临时目录，可以根据实际情况修改
//	router.Static("/upload", "./upload")
//
//	router.POST("/upload", handleUpload)
//
//	router.Run(":8080")
//}
//
//func handleUpload(c *gin.Context) {
//	file, header, err := c.Request.FormFile("file")
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	defer file.Close()
//
//	// 可以根据实际需求修改保存文件的目录
//	uploadDir := "./upload"
//	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
//		os.Mkdir(uploadDir, 0755)
//	}
//
//	// 生成文件名，保留原始文件的扩展名
//	filename := header.Filename
//	outPath := filepath.Join(uploadDir, filename)
//
//	out, err := os.Create(outPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer out.Close()
//
//	_, err = io.Copy(out, file)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File %s uploaded successfully", filename)})
//}