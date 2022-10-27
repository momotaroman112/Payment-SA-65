package controller

import (
	"net/http"

	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
)

// GET /bookings
// List all bookings
func ListBookings(c *gin.Context) {
	var bookings []entity.Booking
	if err := entity.DB().Raw("SELECT * FROM bookings").Preload("Member").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GET /booking/:id
// Get booking by id
func GetBooking(c *gin.Context) {
	var booking entity.Booking
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM bookings WHERE id = ?", id).Preload("Member").Find(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// // POST /bookings
// func CreateBooking(c *gin.Context) {
// 	var booking entity.Booking
// 	if err := c.ShouldBindJSON(&booking); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := entity.DB().Create(&booking).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": booking})
// }

// // PATCH /bookings
// func UpdateBooking(c *gin.Context) {
// 	var booking entity.Booking
// 	if err := c.ShouldBindJSON(&booking); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entity.DB().Where("id = ?", booking.ID).First(&booking); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
// 		return
// 	}

// 	if err := entity.DB().Save(&booking).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": booking})
// }

// // DELETE /bookings/:id
// func DeleteBooking(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM bookings WHERE id = ?", id); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		return
// 	}
// 	/*
// 		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}*/

// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }