package bpjscontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var bpjs []models.Bpjs
	models.DB.Preload("Karyawan.Unit").Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Find(&bpjs)
	c.JSON(http.StatusOK, gin.H{"data": bpjs})
}
func Create(c *gin.Context) {
	var bpjs models.Bpjs

	if err := c.ShouldBindJSON(&bpjs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&bpjs)
	c.JSON(http.StatusOK, gin.H{"data": bpjs})
}
func Show(c *gin.Context) {
	var bpjs models.Bpjs
	id := c.Param("id")
	if err := models.DB.First(&bpjs, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": bpjs})
}
func ShowUnit(c *gin.Context) {
	var bpjs []models.Bpjs
	id := c.Param("id")
	if err := models.DB.Preload("Karyawan.Unit").Preload("Karyawan.Bagian").Preload("Karyawan.Jabatan").Where("unit_id = ?", id).Find(&bpjs).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": bpjs})
}
func Update(c *gin.Context) {
	var bpjs models.Bpjs
	id := c.Param("id")

	if err := c.ShouldBindJSON(&bpjs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&bpjs).Where("id = ?", id).Updates(&bpjs).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var bpjs models.Bpjs
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}
	// if err := c.ShouldBindJSON(&bpjs); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&bpjs, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
