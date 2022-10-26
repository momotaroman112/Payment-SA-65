package controller

import (
	"net/http"

	"github.com/AKiRA3563/sa-65/entity"
	"github.com/gin-gonic/gin"
)

// GET /equipment
func GetEquipment(c *gin.Context) {
	id := c.Param("id")
	var equipment entity.Equipment
	if err := entity.DB().Table("equipment").Where("id = ?", id).Find(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": equipment})
}

// GET /equipment/:id
func ListEquipment(c *gin.Context) {
	var equipment []entity.Equipment
	if err := entity.DB().Table("equipment").Find(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": equipment})
}