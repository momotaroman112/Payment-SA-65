package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Siriwan38/Sa-65-Group-18/entity"
)

func ListUsage(c *gin.Context) {
	var Usage []entity.Usage
	if err := entity.DB().Table("usages").Find(&Usage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Usage})
}
