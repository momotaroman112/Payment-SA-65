package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Siriwan38/Sa-65-Group-18/entity"
)

// POST /Booking
func CreateBooking(c *gin.Context) {

	var User entity.User
	var Room entity.Room
	var Usage entity.Usage
	var Booking entity.Booking

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร Booking
	if err := c.ShouldBindJSON(&Booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//10: ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", Booking.MemberID).First(&User); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// 11: ค้นหา Room ด้วย id
	if tx := entity.DB().Where("id = ?", Booking.RoomID).First(&Room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	// 12: ค้นหา Usage ด้วย id
	if tx := entity.DB().Where("id = ?", Booking.UsageID).First(&Usage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	// 14: สร้าง Booking
	lr := entity.Booking{
		Member:           User,
		Room:             Room,
		Usage:            Usage,
		BookingTimeStart: Booking.BookingTimeStart,
		BookingTimeStop:  Booking.BookingTimeStop,
	}

	// 15: บันทึก
	if err := entity.DB().Create(&lr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lr})
}

// GET: /api/
func ListBooking(c *gin.Context) {
	var Booking []*entity.Booking
	if err := entity.DB().Table("bookings").Preload("Room").Preload("Usage").Preload("Member").Find(&Booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Booking})
}
