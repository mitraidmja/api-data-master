package absensicontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var absensi []models.Absensi
	models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Find(&absensi)
	c.JSON(http.StatusOK, gin.H{"data": absensi})
}
func Show(c *gin.Context) {
	var absensi []models.Absensi
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").First(&absensi, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": absensi})
}
func ShowUnit(c *gin.Context) {
	var absensi []models.Absensi
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("unit_id = ?", id).Find(&absensi).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": absensi})
}
func FilterDate(c *gin.Context) {
	var absensi []models.Absensi
	unitid := c.Query("unitid")
	startdate := c.Query("sdate")
	enddate := c.Query("edate")
	if err := models.DB.Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Preload("Karyawan.Unit").Where("unit_id = ? AND tgl BETWEEN ? AND ?", unitid, startdate, enddate).Find(&absensi).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": absensi})
}
func Create(c *gin.Context) {
	var absensi models.Absensi
	if err := c.ShouldBindJSON(&absensi); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&absensi)
	c.JSON(http.StatusOK, gin.H{"data": absensi})

}
func Update(c *gin.Context) {
	var absensi models.Absensi
	id := c.Param("id")

	if err := c.ShouldBindJSON(&absensi); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&absensi).Where("id = ?", id).Updates(&absensi).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var absensi models.Absensi
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}

	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&absensi, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
