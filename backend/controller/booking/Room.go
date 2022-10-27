package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Siriwan38/Sa-65-Group-18/entity"
)

func ListRoom(c *gin.Context) {
	var Room []entity.Room
	if err := entity.DB().Table("rooms").Find(&Room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Room})
}
