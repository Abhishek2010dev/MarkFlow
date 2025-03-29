package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Abhishek2010dev/MarkFlow/utils"
	"github.com/gin-gonic/gin"
)

func UploadMarkdown(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if filepath.Ext(file.Filename) != ".md" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file type, only .md files are allowed",
		})
		return
	}

	filename := filepath.Base(file.Filename)
	if utils.FileExists(filename) {
		c.JSON(http.StatusConflict, gin.H{
			"error": "File already exists",
		})
		return
	}

	savePath := filepath.Join("./uploads", filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload file",
		})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", filename))
}

