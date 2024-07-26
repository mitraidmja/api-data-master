package unitcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	// "github.com/adewputro/go-rest-api-mja/select"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var units []models.Unit
	models.DB.Preload("Mou").Preload("UnitStaff").Preload("UnitStaff.Staff").Preload("Karyawan", "is_aktif = ?", 0).Find(&units)

	c.JSON(http.StatusOK, gin.H{"data": units})
}
func SelectData(c *gin.Context) {

	var units []models.Result
	models.DB.Select("unit").Table("units").Scan(&units)

	c.JSON(http.StatusOK, gin.H{"data": units})
}
func Show(c *gin.Context) {
	var units []models.Unit
	id := c.Param("id")

	if err := models.DB.Preload("Mou").Preload("UnitStaff").Preload("UnitStaff.Staff").Preload("Karyawan", "is_aktif = ?", 0).Find(&units, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": units})

}
func ShowGroup(c *gin.Context) {
	var units []models.Unit
	id := c.Param("id")
	if err := models.DB.Preload("Mou").Preload("Staff").Preload("Karyawan").Where("groups = ?", id).Find(&units).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": units})

}

func Create(c *gin.Context) {

	var units models.Unit
	if err := c.ShouldBindJSON(&units); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&units)

	c.JSON(http.StatusOK, gin.H{"data": units})

}
func Update(c *gin.Context) {
	var units models.Unit
	id := c.Param("id")

	if err := c.ShouldBindJSON(&units); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&units).Where("id = ?", id).Updates(&units).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var units models.Unit
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}

	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&units, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
