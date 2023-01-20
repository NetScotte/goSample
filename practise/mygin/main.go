package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func jsonHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func htmlHandler(c *gin.Context) {
	// gin.H = map[string]interface{}
	//contents := &gin.H{
	//	"content": []string{"this", "is", "from", "db"},
	//}
	// list
	contents := []string{"this", "is", "from", "db"}

	c.HTML(http.StatusOK, "index.html", contents)
}

func postHandler(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("parse data error: %v\n", err)
	} else {
		fmt.Printf("%s\n", data)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func main() {
	// 设置模式gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 加载所有文件，多级目录可以使用views/**/*
	router.LoadHTMLGlob("practise/mygin/views/*")
	// 加载指定的文件
	// router.LoadHTMLFiles("views/index.html")
	router.GET("/json", jsonHandler)
	router.GET("/html", htmlHandler)
	router.POST("/post", postHandler)
	// 默认监听0.0.0.0:8080
	router.Run("localhost:9191")
}
