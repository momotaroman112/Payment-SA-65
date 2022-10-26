package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siriwan38/sa-65-example/entity"
)

func ListUsage(c *gin.Context) {
	var Usage []entity.Usage
	if err := entity.DB().Table("usages").Find(&Usage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Usage})
}
