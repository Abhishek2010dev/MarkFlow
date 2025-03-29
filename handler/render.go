package handler

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Abhishek2010dev/MarkFlow/utils"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

func GetMarkdownContentAsHtml(c *gin.Context) {
	rawFilename := c.Param("filename")
	filename := filepath.Base(rawFilename)

	if filepath.Ext(filename) != ".md" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file type, only .md files are allowed",
		})
		return
	}

	if !utils.FileExists(filename) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "File does not exist",
		})
		return
	}

	filePath := filepath.Join("./uploads", filename)
	content, err := utils.ReadFileContent(filePath)
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

