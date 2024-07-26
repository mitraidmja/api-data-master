package karyawancontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var karyawans []models.Karyawan
	models.DB.Preload("Unit").Preload("Bagian").Preload("Jabatan").Where("is_aktif = 0").Find(&karyawans)
	c.JSON(http.StatusOK, gin.H{"data": karyawans})

}
func Show(c *gin.Context) {
	var karyawans models.Karyawan
	id := c.Param("id")

	if err := models.DB.Preload("Unit").Preload("Bagian").Preload("Jabatan").First(&karyawans, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}

	c.JSON(http.StatusOK, gin.H{"data": karyawans})

}
func ShowUnit(c *gin.Context) {
	var karyawans []models.Karyawan
	id := c.Param("id")

	if err := models.DB.Preload("Unit").Preload("Bagian").Preload("Jabatan").Where("unit_id = ? AND is_aktif = 0", id).Find(&karyawans).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}

	c.JSON(http.StatusOK, gin.H{"data": karyawans})

}

func Create(c *gin.Context) {
	var karyawan models.Karyawan

	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"data": karyawan})
}
func Update(c *gin.Context) {
	var karyawan models.Karyawan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&karyawan).Where("id = ?", id).Updates(&karyawan).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var karyawan models.Karyawan
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&karyawan, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
