package controller

import (
	"net/http"

	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
)

// GET /bookings----> list
func ListBooking(c *gin.Context) {
	var bookings []entity.Booking
	if err := entity.DB().Raw("SELECT * FROM Bookings").
	Preload("Member").
	Preload("Room").
	Preload("Room.Type").
	Preload("FoodOrdereds").
	Preload("FoodOrdereds.FoodOrderedFoodSets").
	Preload("FoodOrdereds.FoodOrderedFoodSets.FoodSet").
	Find(&bookings).Error; err != nil {
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
	Preload("Member").
	Preload("Room").
	Preload("Room.Type").
	Preload("FoodOrdereds").
	Preload("FoodOrdereds.FoodOrderedFoodSets").
	Preload("FoodOrdereds.FoodOrderedFoodSets.FoodSet").
	Find(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// GET /booking/user/:id
func GetBookingByUser(c *gin.Context) {
	id := c.Param("id")
	var booking entity.Booking
	if err := entity.DB().Raw("SELECT * FROM bookings WHERE member_id = ?", id).
	Preload("Member").
	Preload("Room").
	Preload("Room.Type").
	Preload("FoodOrdereds").
	Preload("FoodOrdereds.FoodOrderedFoodSets").
	Preload("FoodOrdereds.FoodOrderedFoodSets.FoodSet").
	Find(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}
