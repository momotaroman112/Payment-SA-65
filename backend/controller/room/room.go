package controller

import (
	"net/http"

	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {

	var room entity.Room
	var employee entity.Employee
	var ttype entity.Type
	var building entity.Building
	var serviceday entity.ServiceDay
	var period entity.Period

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", room.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
	}

	// 11: ค้นหา type ด้วย id
	if tx := entity.DB().Where("id = ?", room.TypeID).First(&ttype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room type not found"})
	}

	// 12: ค้นหา building ด้วย id
	if tx := entity.DB().Where("id = ?", room.BuildingID).First(&building); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "building type not found"})
	}

	// 13: ค้นหา serviceday ด้วย id
	if tx := entity.DB().Where("id = ?", room.ServiceDayID).First(&serviceday); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Service day not found"})
	}

	// 14: ค้นหา period ด้วย id
	if tx := entity.DB().Where("id = ?", room.PeriodID).First(&period); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "period not found"})
	}

	// 15: สร้าง Room
	WV := entity.Room{	
		Number: 	room.Number,	// ตั้งค่าฟิลด์ RoomName
		Employee:   employee,   	// โยงความสัมพันธ์กับ Entity employee
		Type:       ttype,      // โยงความสัมพันธ์กับ Entity Type
		Building:   building,   // โยงความสัมพันธ์กับ Entity Building
		ServiceDay: serviceday, // โยงความสัมพันธ์กับ Entity ServiceDay
		Period:     period,     // โยงความสัมพันธ์กับ Entity Period
		Name:       room.Name,  // ตั้งค่าฟิลด์ Name
	}

	// 16: บันทึก
	if err := entity.DB().Create(&WV).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": WV})
}

// GET /employee
func ListRoom(c *gin.Context) {
	var room []entity.Room
	if err := entity.DB().Table("rooms").Preload("Employee").Preload("Type").Preload("Building").Preload("ServiceDay").Preload("Period").Find(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room})
}
