package controller

import (
	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListServiceDay(c *gin.Context){
	var ServiceDay []entity.ServiceDay
	if err := entity.DB().Table("service_days").Find(&ServiceDay).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"data": ServiceDay})
}