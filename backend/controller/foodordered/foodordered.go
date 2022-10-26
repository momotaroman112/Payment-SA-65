package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /foodordereds
func CreateFoodOrdered(c *gin.Context) {

	var foodordered entity.FoodOrdered
	var booking entity.Booking
	// var foodset entity.FoodSet
	var foodpayment_type entity.FoodPaymentType

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร foodordered
	if err := c.ShouldBindJSON(&foodordered); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา booking ด้วย id
	if tx := entity.DB().Where("id = ?", foodordered.BookingID).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	// 10: ค้นหา foodset ด้วย id
	for _, orderFoodSet := range foodordered.FoodOrderedFoodSets {
		if tx := entity.DB().Where("id = ?", orderFoodSet.FoodSetID).First(&orderFoodSet); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "food set not found"})
			return
		}
	}

	// 11: ค้นหา foodpayment_type ด้วย id
	if tx := entity.DB().Where("id = ?", foodordered.FoodPaymentTypeID).First(&foodpayment_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food payment type not found"})
		return
	}
	// 12: สร้าง FoodOrdered
	wv := entity.FoodOrdered{
		Booking:             booking,                         // โยงความสัมพันธ์กับ Entity Booking
		FoodPaymentType:     foodpayment_type,                // โยงความสัมพันธ์กับ Entity FoodPaymentType
		FoodTime:            foodordered.FoodTime,            // ตั้งค่าฟิลด์ FoodTime
		FoodOrderedFoodSets: foodordered.FoodOrderedFoodSets, // โยงความสัมพันธ์กับ Entity FoodSet (แต่ไม่โดยตรง เพราะเป็นคสพแบบหลาย)
		TotalPrice:          foodordered.TotalPrice,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /foodordereds
func ListFoodOrdereds(c *gin.Context) {
	var foodOrdered []entity.FoodOrdered

	/*เงื่อนไขสำหรับการค้นหา โดยดึงข้อมูลจากตารางรองที่เกี่ยวข้องมา #ระวัง ชื่อ field ต้องตรงกัน
	ซึ่งดูฟิลด์ได้จากเราสร้างไว้ให้ entity หลัก ในไฟล์ schema */

	if err := entity.DB().Raw("SELECT * FROM food_ordereds").
		Preload("Booking").Preload("FoodPaymentType").
		Preload("FoodOrderedFoodSets").Preload("FoodOrderedFoodSets.FoodSet"). //preload แบบ join table
		Find(&foodOrdered).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foodOrdered})

}

// GET /foodordered/:id
func GetFoodOrdered(c *gin.Context) {
	var foodOrdered entity.FoodOrdered //GET จะ​ get มาแค่ก้อนเดียวเลยไม่ใช้ array (เก็ทไอดีของตัวที่เคยบันทึก) [ex. เก็ทเอาไปคิดราคา(ของระบบอื่น)]

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM food_ordereds WHERE id = ?", id).
		Preload("Booking").Preload("FoodPaymentType").
		Preload("FoodOrderedFoodSets").Preload("FoodOrderedFoodSets.FoodSet").
		Find(&foodOrdered).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foodOrdered})

}

/* Note.
1. ทำไม foodpayment_type ถึงยังมี underscore ? ลืมแก้รึป่าว
2. ทำความเข้าใจ for loop ตรง FoodSet
3. ตรง func CreateFoodOrdered ทำไมต้องคอมเมนต์ foodset ไว้หรอ
4. ใช้ command + D จ่ะ จะได้ไวๆ
*/
