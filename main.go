package main

import (
	"github.com/Abhishek2010dev/MarkFlow/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/check", handler.CheckGrammerHandler)
	router.POST("/upload", handler.UploadMarkdownHandler)
	router.Run()
}
