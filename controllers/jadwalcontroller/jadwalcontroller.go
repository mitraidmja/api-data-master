package jadwalcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	// "github.com/adewputro/go-rest-api-mja/select"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func Index(c *gin.Context) {
	var jadwals []models.Jadwal
	models.DB.Find(&jadwals)
	c.JSON(http.StatusOK, gin.H{"data": jadwals})
}

func Show(c *gin.Context) {
	var jadwals []models.Jadwal
	id := c.Param("id")
	models.DB.Where("unit_id = ?", id).Find(&jadwals)
	c.JSON(http.StatusOK, gin.H{"data": jadwals})
}

func Create(c *gin.Context) {
	var jadwals models.Jadwal
	if err := c.ShouldBindJSON(&jadwals); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&jadwals)
	c.JSON(http.StatusOK, gin.H{"data": jadwals})
}

func Update(c *gin.Context) {
	var jadwals models.Jadwal
	id := c.Param("id")
	if err := c.ShouldBindJSON(&jadwals); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&jadwals).Where("id = ?", id).Updates(&jadwals).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}

func Delete(c *gin.Context) {
	var jadwals models.Jadwal
	id := c.Param("id")
	if models.DB.Delete(&jadwals, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
