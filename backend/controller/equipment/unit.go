package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Siriwan38/Sa-65-Group-18/entity"
)
// POST /units
func CreateUnit(c *gin.Context) {
	var unit entity.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&unit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": unit})
}
// GET /unit/:id
func GetUnit(c *gin.Context) {
	var unit entity.Unit
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&unit); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unit not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": unit})
}
// GET /units
func ListUnits(c *gin.Context) {
	var units []entity.Unit
	if err := entity.DB().Raw("SELECT * FROM units").Scan(&units).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": units})
}
// DELETE /units/:id
func DeleteUnit(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM units WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unit not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
// PATCH /units
func UpdateUnit(c *gin.Context) {
	var unit entity.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", unit.ID).First(&unit); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unit not found"})
		return
	}
	if err := entity.DB().Save(&unit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": unit})
}