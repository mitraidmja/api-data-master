package kontrakcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var kontraks []models.Kontrak
	models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Find(&kontraks)
	c.JSON(http.StatusOK, gin.H{"data": kontraks})
}
func Show(c *gin.Context) {
	var kontraks models.Kontrak
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").First(&kontraks, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": kontraks})
}
func ShowKaryawan(c *gin.Context) {
	var kontraks []models.Kontrak
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("karyawan_id = ?", id).Find(&kontraks).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": kontraks})
}
func ShowUnit(c *gin.Context) {
	var kontraks []models.Kontrak
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("unit_id = ?", id).Find(&kontraks).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": kontraks})
}
func Create(c *gin.Context) {
	var kontraks models.Kontrak

	if err := c.ShouldBindJSON(&kontraks); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&kontraks)
	c.JSON(http.StatusOK, gin.H{"data": kontraks})
}
func Update(c *gin.Context) {
	var kontraks models.Kontrak
	id := c.Param("id")

	if err := c.ShouldBindJSON(&kontraks); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&kontraks).Where("id = ?", id).Updates(&kontraks).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var kontraks models.Kontrak
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&kontraks, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
