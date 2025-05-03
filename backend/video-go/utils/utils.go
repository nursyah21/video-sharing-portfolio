package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindBodyWithJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "must have body"})
		return false
	}

	return true
}

func CheckExist(c *gin.Context, db *gorm.DB, msg string, dest any, conds ...any) bool {
	if err := db.First(&dest, conds...).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return false
	}
	return true
}

func DeleteIfExist(c *gin.Context, db *gorm.DB, msg string, dest any, conds ...any) bool {
	if err := db.Delete(&dest, conds...).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return false
	}
	return true
}
