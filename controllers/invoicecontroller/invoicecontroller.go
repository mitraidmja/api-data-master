package invoicecontroller

import (
	"net/http"

	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var invoice []models.Invoice
	models.DB.Preload("Unit").Find(&invoice)
	c.JSON(http.StatusOK, gin.H{"data": invoice})
}
func ShowStatus(c *gin.Context) {
	var status []models.Status
	id := c.Param("id")
	if err := models.DB.Preload("Invoice").Preload("Invoice.Unit").Where("invoice_id = ? ", id).Find(&status).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": status})
}
func ShowUnit(c *gin.Context) {
	var invoice []models.Invoice
	id := c.Param("id")
	if err := models.DB.Preload("Unit").Where("unit_id = ? ", id).Find(&invoice).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": invoice})

}
func CreateStatus(c *gin.Context) {
	var status models.Status
	if err := c.ShouldBindJSON(&status); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&status)
	c.JSON(http.StatusOK, gin.H{"data": status})
}
func Create(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&invoice)
	c.JSON(http.StatusOK, gin.H{"data": invoice})
}
func Update(c *gin.Context) {
	var invoice models.Invoice
	id := c.Param("id")

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&invoice).Where("id = ?", id).Updates(&invoice).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	// models.DB.Create(&invoice)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di perbarui"})
}
func Delete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var invoice models.Invoice
	id := c.Param("id")
	// input := map[string]string{"id" : "0"}
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&invoice, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
