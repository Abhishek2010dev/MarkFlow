package handler

import (
	"fmt"
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
			"error": "File type must be Markdown.",
		})
		return
	}

	if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload file",
		})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
