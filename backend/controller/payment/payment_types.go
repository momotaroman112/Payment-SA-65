package controller

import (
	"net/http"

	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
)

// GET /Payment Types  //การส่งข้อมูล
func GetPaymentType(c *gin.Context) {
	var paymentType []entity.PaymentType
	if err := entity.DB().Raw("SELECT * FROM payment_types").Scan(&paymentType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentType})
}

