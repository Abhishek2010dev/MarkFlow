package main

import (
	"github.com/Abhishek2010dev/MarkFlow/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/check-grammer", handler.CheckGrammerHandler)
	router.POST("/notes", handler.UploadMarkdownHandler)
	router.Run()
}
