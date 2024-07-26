package unitstaffcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	// "github.com/adewputro/go-rest-api-mja/select"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var unitstaff []models.UnitStaff
	models.DB.Preload("Unit").Find(&unitstaff)
	c.JSON(http.StatusOK, gin.H{"data": unitstaff})
}

func Show(c *gin.Context) {
	var unitstaff []models.UnitStaff
	id := c.Param("id")
	models.DB.Preload("Unit").Where("kode_staff = ?", id).Find(&unitstaff)
	c.JSON(http.StatusOK, gin.H{"data": unitstaff})
}

func Create(c *gin.Context) {
	var unitstaffs models.UnitStaff
	if err := c.ShouldBindJSON(&unitstaffs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&unitstaffs)

	c.JSON(http.StatusOK, gin.H{"data": unitstaffs})
}
func Delete(c *gin.Context) {
	var unitstaffs models.UnitStaff
	id := c.Param("id")
	if models.DB.Where("staff_id = ?", id).Delete(&unitstaffs).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
