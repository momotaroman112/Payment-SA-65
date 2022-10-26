package controller

import (
	"net/http"

	"github.com/momotaroman112/sa-65-file/entity"
	"github.com/gin-gonic/gin"
)

// GET /FoodOrdereds
func GetFoodOrdered(c *gin.Context) {
	var foodordered []entity.FoodOrdered
	if err := entity.DB().Raw("SELECT * FROM food_ordereds").Scan(&foodordered).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodordered})
}
