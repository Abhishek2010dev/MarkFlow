package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckGrammerHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"filename": file.Filename,
	})
}
