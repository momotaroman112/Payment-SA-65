package controller

import (
	"net/http"

	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
)

/* Get คือดึงตาม id ที่ส่งไป(ส่งไปหรือส่งมาว้ะ 5555) ส่วน list คือดึงทั้งหมด*/


// POST /foodpayment_types
func CreateFoodPaymentType(c *gin.Context) {
	var foodpayment_type entity.FoodPaymentType
	if err := c.ShouldBindJSON(&foodpayment_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&foodpayment_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foodpayment_type})
}

// GET /foodpayment_type/:id
func GetFoodPaymentType(c *gin.Context) {
	var foodpayment_type entity.FoodPaymentType

	//ใช้ Preload("Owner") หรอ?
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM food_payment_types WHERE id = ?", id).Find(&foodpayment_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodpayment_type})
}

// GET /foodpayment_types
func ListFoodPaymentTypes(c *gin.Context) {
	var foodpayment_type []entity.FoodPaymentType

	if err := entity.DB().Raw("SELECT * FROM food_payment_types").Scan(&foodpayment_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodpayment_type})
}

// DELETE /foodpayment_types/:id
// func DeleteFoodPaymentType(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM food_payment_types WHERE id = ?", id); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }


// PATCH /foodpayment_types
// func UpdateFoodPaymentType(c *gin.Context) {
// 	var foodpayment_type entity.FoodPaymentType
// 	if err := c.ShouldBindJSON(&foodpayment_type); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entity.DB().Where("id = ?", foodpayment_type.ID).First(&foodpayment_type); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
// 		return
// 	}

// 	if err := entity.DB().Save(&foodpayment_type).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": foodpayment_type})
// }

/* Note.
1. preload จะใช้ก็ต่อเมื่อ ตารางหลักต้องการดึงรายละเอียดต่างๆ(ข้อมูล)ของตารางรองไปใส่
ก็คือปกติมันจะดึงแค่ตัว ID ไปแต่มันจะไม่ดึงพวก object เราเขียน preload เพื่อดึงรายละเอียด
มันไปด้วย ex. ตาราง Listbookings ที่มีการ preload("member")
2. เมื่อเราใส่ preload จะใส่ scan ไม่ได้ ใส่ได้แค่ find
3. ชื่อที่ใช้ใน database มันจะใช้เป็นพหูพจน์ (เติม s)
4.
*/
