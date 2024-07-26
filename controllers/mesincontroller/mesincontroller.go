package mesincontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var mesins []models.Mesin
	models.DB.Preload("Unit").Find(&mesins)
	c.JSON(http.StatusOK, gin.H{"data": mesins})

}
func Create(c *gin.Context) {

	var mesins models.Mesin
	if err := c.ShouldBindJSON(&mesins); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&mesins)

	c.JSON(http.StatusOK, gin.H{"data": mesins})

}
func Show(c *gin.Context) {
	var mesins []models.Mesin
	id := c.Param("id")

	if err := models.DB.First(&mesins, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": mesins})

}
func ShowUnit(c *gin.Context) {
	var mesins []models.Mesin
	id := c.Param("id")
	if err := models.DB.Where("unit_id = ?", id).Find(&mesins).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": mesins})
}
func Update(c *gin.Context) {
	var mesins models.Mesin
	id := c.Param("id")

	if err := c.ShouldBindJSON(&mesins); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&mesins).Where("id = ?", id).Updates(&mesins).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var mesins models.Mesin
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}

	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&mesins, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
