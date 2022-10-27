package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Siriwan38/Sa-65-Group-18/entity"
)

func ListMember(c *gin.Context) {
	var User []entity.User
	if err := entity.DB().Table("users").Find(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": User})
}
func GetMember(c *gin.Context) {
	id := c.Param("id")
	var User entity.User
	if err := entity.DB().Table("users").Where("id = ?", id).Find(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": User})
}
