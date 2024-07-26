package logincontroller

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"github.com/adewputro/go-rest-api-mja/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var body struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Print(body.Username)
	var karyawans models.Staff
	// id := c.Param("id")

	if err := models.DB.Where("username = ?", body.Username).First(&karyawans).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak di Temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

	}
	// hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	err := bcrypt.CompareHashAndPassword([]byte(karyawans.Password), []byte(body.Password))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Password Salah"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": karyawans.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SCREET")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Token null"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user" : &karyawans})

}
