package groupcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var groupusers []models.GroupUser
	models.DB.Find(&groupusers)
	c.JSON(http.StatusOK, gin.H{"data": groupusers})
}

func Create(c *gin.Context) {

	var groupusers models.GroupUser
	if err := c.ShouldBindJSON(&groupusers); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&groupusers)

	c.JSON(http.StatusOK, gin.H{"data": groupusers})

}

func Show(c *gin.Context) {
	var groupusers []models.GroupUser
	id := c.Param("id")

	if err := models.DB.First(&groupusers, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": groupusers})

}

func Update(c *gin.Context) {
	var groupusers models.GroupUser
	id := c.Param("id")

	if err := c.ShouldBindJSON(&groupusers); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&groupusers).Where("id = ?", id).Updates(&groupusers).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&groupusers)

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}

func Delete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var groupusers models.GroupUser
	id := c.Param("id")
	if models.DB.Delete(&groupusers, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
