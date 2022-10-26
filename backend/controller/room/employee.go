package controller

import (
	"github.com/PimchanokS/sa-64-example/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /employee/:id
func GetEmployee(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if err := entity.DB().Table("employees").Where("id = ?",id).Find(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /employee
// List all employees
func ListEmployee(c *gin.Context) {
	var employee []entity.Employee
	if err := entity.DB().Table("employees").Find(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}


