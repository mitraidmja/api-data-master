package staffcontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var staffs []models.Staff
	models.DB.Preload("UnitStaff").Preload("UnitStaff.Unit").Find(&staffs)
	c.JSON(http.StatusOK, gin.H{"data": staffs})

}
func Show(c *gin.Context) {
	var staffs models.Staff
	id := c.Param("id")

	if err := models.DB.Preload("UnitStaff").First(&staffs, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}

	c.JSON(http.StatusOK, gin.H{"data": staffs})

}
func ShowUnit(c *gin.Context) {
	var staffs []models.Staff
	id := c.Param("id")

	if err := models.DB.Where("unit_id = ?", id).Find(&staffs).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}

	c.JSON(http.StatusOK, gin.H{"data": staffs})

}

func Create(c *gin.Context) {
	var staffs models.Staff
	if err := c.ShouldBindJSON(&staffs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&staffs)
	c.JSON(http.StatusOK, gin.H{"data": staffs})
}
func Update(c *gin.Context) {
	var staffs models.Staff
	id := c.Param("id")

	if err := c.ShouldBindJSON(&staffs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&staffs).Where("id = ?", id).Updates(&staffs).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&Staff)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var staffs models.Staff
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&staffs, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
