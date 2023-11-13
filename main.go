package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/uploads", "./uploads")
	router.POST("/upload", func(c *gin.Context) {
		lat := c.PostForm("latitude")
		long := c.PostForm("longitude")
		latitude, _ := strconv.ParseFloat(lat, 64)
		longitude, _ := strconv.ParseFloat(long, 64)
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		// Tạo thư mục nếu nó chưa tồn tại
		err = os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			c.JSON(400, gin.H{"message": "Failed to create directory"})
			return
		}
		saveFilename := "uploads/" + file.Filename
		// Lưu file vào thư mục uploads
		err = c.SaveUploadedFile(file, saveFilename)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to save file"})
			return
		}
		fmt.Printf("latitude: %v\n", latitude)
		fmt.Printf("longitude: %v\n", longitude)
		SetGeotag(saveFilename, latitude, longitude)

		c.File(saveFilename)
	})

	router.Run(":8080")
}
