package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AKiRA3563/sa-65/entity"
)

// GET /employee
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee entity.Employee
	if err := entity.DB().Table("employees").Where("id = ?", id).Find(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /employee/:id
func ListEmployee(c *gin.Context) {
	var employee []entity.Employee
	if err := entity.DB().Table("employees").Find(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}