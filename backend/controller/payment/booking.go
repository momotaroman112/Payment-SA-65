package controller

import (
	"net/http"

	"github.com/momotaroman112/sa-65-file/entity"
	"github.com/gin-gonic/gin"
)

// GET /bookings----> list
func ListBooking(c *gin.Context) {
	var bookings []entity.Booking
	if err := entity.DB().Raw("SELECT * FROM Bookings").
	Preload("FoodOrdereds").
	Preload("FoodOrdereds.FoodOrderedFoodSets").
	Preload("FoodOrdereds.FoodOrderedFoodSets.FoodSet").
	Preload("User").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GET /booking/:id
func GetBooking(c *gin.Context) {
	id := c.Param("id")
	var booking entity.Booking
	if err := entity.DB().Raw("SELECT * FROM bookings WHERE id = ? ORDER BY room ASC", id).
	Preload("FoodOrdereds").
	Preload("FoodOrdereds.FoodOrderedFoodSets").
	Preload("FoodOrdereds.FoodOrderedFoodSets.FoodSet").
	Preload("User").Find(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// GET /booking/user/:id
func GetBookingByUser(c *gin.Context) {
	id := c.Param("id")
	var booking entity.Booking
	if err := entity.DB().Raw("SELECT * FROM bookings WHERE user_id = ? ORDER BY room ASC", id).Preload("User").Find(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}
