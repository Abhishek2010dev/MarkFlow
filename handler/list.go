package handler

import (
	"net/http"

	"github.com/Abhishek2010dev/MarkFlow/utils"
	"github.com/gin-gonic/gin"
)

func ListFilesHandler(c *gin.Context) {
	dirPath := "./uploads/"

	files, err := utils.ListFiles(dirPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
