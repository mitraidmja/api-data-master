package kecelakaancontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	// "github.com/adewputro/go-rest-api-mja/select"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func Index(c *gin.Context) {
	var kecelakaans []models.Kecelakaan
	models.DB.Find(&kecelakaans)
	c.JSON(http.StatusOK, gin.H{"data": kecelakaans})
}
func Show(c *gin.Context) {
	var kecelakaans []models.Kecelakaan
	id := c.Param("id")
	models.DB.Where("unit_id = ?", id).Find(&kecelakaans)
	c.JSON(http.StatusOK, gin.H{"data": kecelakaans})
}
func Create(c *gin.Context) {
	var kecelakaans models.Kecelakaan
	if err := c.ShouldBindJSON(&kecelakaans); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&kecelakaans)
	c.JSON(http.StatusOK, gin.H{"data": kecelakaans})
}
func Update(c *gin.Context) {
	var kecelakaans models.Kecelakaan
	id := c.Param("id")
	if err := c.ShouldBindJSON(&kecelakaans); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&kecelakaans).Where("id = ?", id).Updates(&kecelakaans).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var kecelakaans models.Kecelakaan
	id := c.Param("id")
	if models.DB.Delete(&kecelakaans, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
