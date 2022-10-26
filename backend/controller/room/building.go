package controller

import (
	"github.com/PimchanokS/sa-64-example/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)


func ListBuilding(c *gin.Context){
	var Building []entity.Building
	if err := entity.DB().Table("buildings").Find(&Building).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"data": Building})
}
