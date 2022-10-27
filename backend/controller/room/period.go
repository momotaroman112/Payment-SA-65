package controller

import (
	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListPeriod(c *gin.Context){
	var Period []entity.Period
	if err := entity.DB().Table("periods").Find(&Period).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"data": Period})
}