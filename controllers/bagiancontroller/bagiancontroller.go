package bagiancontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var bagians []models.Bagian
	models.DB.Find(&bagians)
	c.JSON(http.StatusOK, gin.H{"data": bagians})
}

func Show(c *gin.Context) {
	var bagians []models.Bagian
	id := c.Param("id")
	if err := models.DB.First(&bagians, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak ada Data"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": bagians})
}
func Create(c *gin.Context) {
	var bagians models.Bagian
	if err := c.ShouldBindJSON(&bagians); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&bagians)
	c.JSON(http.StatusOK, gin.H{"data": bagians})

}
func Update(c *gin.Context) {
	var bagians models.Bagian
	id := c.Param("id")

	if err := c.ShouldBindJSON(&bagians); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&bagians).Where("id = ?", id).Updates(&bagians).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	var bagian models.Bagian
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}

	
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&bagian, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
