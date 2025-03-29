package handler

import (
	"log"
	"net/http"

	"github.com/Abhishek2010dev/MarkFlow/grammer"
	"github.com/Abhishek2010dev/MarkFlow/utils"
	"github.com/gin-gonic/gin"
)

func CheckGrammerHandler(c *gin.Context) {
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

	f, err := file.Open()
	if err != nil {
		log.Printf("Failed to open file: %s \n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer f.Close()

	content, err := utils.GetFileContent(file)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to access content of file",
		})
		return
	}

	grammerError, err := grammer.CheckGrammer(content)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check grammer",
		})
		return
	}

	c.JSON(http.StatusOK, grammerError)
}
