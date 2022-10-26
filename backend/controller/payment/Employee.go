package controller

import (
	"net/http"

	"github.com/momotaroman112/sa-65-file/entity"
	"github.com/gin-gonic/gin"
)

// GET /employee/:id
// Get Employee by Employee_id
func GetEmployee(c *gin.Context) {
	var Employee entity.Employee
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM employees WHERE id = ?", id).First(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

// GET /employees
// List all Employee
func ListEmployee(c *gin.Context) {
	var Employee []entity.Employee
	if err := entity.DB().Raw("SELECT * FROM employees").Scan(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

// POST /Employees
func CreateEmployee(c *gin.Context) {
	var Employee entity.Employee
	if err := c.ShouldBindJSON(&Employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}
