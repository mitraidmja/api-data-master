package suratperingcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var sp []models.SuratPeringatan
	models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Find(&sp)
	c.JSON(http.StatusOK, gin.H{"data": sp})
}
func Show(c *gin.Context) {
	var sp []models.SuratPeringatan
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").First(&sp, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": sp})
}
func ShowKaryawan(c *gin.Context) {
	var sp []models.SuratPeringatan
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("karyawan_id = ?", id).First(&sp).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": sp})
}
func ShowUnit(c *gin.Context) {
	var sp []models.SuratPeringatan
	id := c.Param("id")
	models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("unit_id = ?", id).First(&sp)
	c.JSON(http.StatusOK, gin.H{"data": sp})
}
func Create(c *gin.Context) {
	var sp models.SuratPeringatan
	if err := c.ShouldBindJSON(&sp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&sp)
	c.JSON(http.StatusOK, gin.H{"data": sp})

}
func Update(c *gin.Context) {
	var sp models.SuratPeringatan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&sp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&sp).Where("id = ?", id).Updates(&sp).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var sp models.SuratPeringatan
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}

	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&sp, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
