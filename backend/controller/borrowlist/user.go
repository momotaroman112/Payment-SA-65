package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AKiRA3563/sa-65/entity"
)

func ListUser(c *gin.Context) {
	var user []entity.User
	if err := entity.DB().Table("users").Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user entity.User
	if err := entity.DB().Table("users").Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}