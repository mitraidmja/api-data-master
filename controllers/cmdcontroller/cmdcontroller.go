package cmdcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var cmdlogs []models.Cmdlog
	models.DB.Find(&cmdlogs)
	c.JSON(http.StatusOK, gin.H{"data": cmdlogs})
}
func Create(c *gin.Context) {
	var cmdlogs models.Cmdlog

	if err := c.ShouldBindJSON(&cmdlogs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&cmdlogs)
	c.JSON(http.StatusOK, gin.H{"data": cmdlogs})
}
