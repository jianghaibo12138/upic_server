package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func UploadImagesController(context *gin.Context) {
	file, err := context.FormFile("image")
	if err != nil {
		context.JSON(200, gin.H{"code": 500, "message": "Got error", "error": err.Error()})
		return
	}

	timeStr := time.Now().Format("20060102150405")

	fileName := fmt.Sprintf("images/%s_%s", timeStr, file.Filename)

	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		context.JSON(200, gin.H{"code": 500, "message": "Got error", "error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"code": 200, "message": "Success", "url": fileName})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.StaticFS("/images", http.Dir("images"))
	r.POST("/upload_images", UploadImagesController)
	r.Run(":8500") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
