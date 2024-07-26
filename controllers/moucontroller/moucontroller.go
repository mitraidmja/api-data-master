package moucontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var mous []models.Mou
	models.DB.Preload("Unit").Preload("Unit.UnitStaff").Preload("Unit.UnitStaff.Staff").Find(&mous)
	c.JSON(http.StatusOK, gin.H{"data": mous})
}

func Show(c *gin.Context) {
	var mous []models.Mou
	id := c.Param("id")
	models.DB.Preload("Unit").Preload("Unit.UnitStaff").Preload("Unit.UnitStaff.Staff").Where("unit_id = ?", id).Find(&mous)
	c.JSON(http.StatusOK, gin.H{"data": &mous})

}

func Create(c *gin.Context) {
	var mous models.Mou
	if err := c.ShouldBindJSON(&mous); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&mous)
	c.JSON(http.StatusOK, gin.H{"data": &mous})
}
func Update(c *gin.Context) {
	var mous models.Mou
	id := c.Param("id")

	if err := c.ShouldBindJSON(&mous); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&mous).Where("id = ?", id).Updates(&mous).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}

func Delete(c *gin.Context) {
	var mous models.Mou
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}

	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&mous, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat dihapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
