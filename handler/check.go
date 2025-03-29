package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Abhishek2010dev/MarkFlow/grammer"
	"github.com/Abhishek2010dev/MarkFlow/utils"
	"github.com/gin-gonic/gin"
)

func CheckGrammer(c *gin.Context) {
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

	f, err := file.Open()
	if err != nil {
		log.Printf("Failed to open file: %s \n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer f.Close()

	content, err := utils.GetFileContent(&f)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to access content of file",
		})
		return
	}

	grammerError, err := grammer.CheckGrammer(string(content))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check grammer",
		})
		return
	}

	c.JSON(http.StatusOK, grammerError)
}
