package handler

import (
	"bytes"
	"log"
	"net/http"

	"github.com/Abhishek2010dev/MarkFlow/utils"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

func GetMarkdownContentAsHtml(c *gin.Context) {
	filename := "./uploads/" + c.Param("filename")
	if !utils.FileExits(filename) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "File does not exist",
		})
		return
	}

	content, err := utils.ReadFileContent(filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read content",
		})
		return
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		log.Printf("Failed to parse markdown: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse markdown",
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
}

