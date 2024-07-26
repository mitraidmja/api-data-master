package keluarcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var keluar []models.KaryawanKeluar
	models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Find(&keluar)
	c.JSON(http.StatusOK, gin.H{"data": keluar})
}
func Show(c *gin.Context) {
	var keluar []models.KaryawanKeluar
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").First(&keluar, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": keluar})
}
func ShowKaryawan(c *gin.Context) {
	var keluar []models.KaryawanKeluar
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("karyawan_id = ?", id).Find(&keluar).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": keluar})
}
func ShowUnit(c *gin.Context) {
	var keluar []models.KaryawanKeluar
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("unit_id = ?", id).Find(&keluar).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": keluar})
}
func Create(c *gin.Context) {
	var keluar models.KaryawanKeluar
	if err := c.ShouldBindJSON(&keluar); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&keluar)
	c.JSON(http.StatusOK, gin.H{"data": keluar})

}
func Update(c *gin.Context) {
	var keluar models.KaryawanKeluar
	id := c.Param("id")

	if err := c.ShouldBindJSON(&keluar); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&keluar).Where("id = ?", id).Updates(&keluar).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var keluar models.KaryawanKeluar
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}

	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&keluar, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
