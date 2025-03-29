package handler

import (
	"net/http"

	"github.com/Abhishek2010dev/MarkFlow/utils"
	"github.com/gin-gonic/gin"
)

func UploadMarkdownHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsValidMarkdown(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "File type must be Markdown.",
		})
		return
	}
}
