package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/momotaroman112/sa-65-file/entity"
)

// POST /bills //รับข้อมูลมาจาก Frontend มาบันทึกลง DB
func Createbill(c *gin.Context) {

	var bill entity.Bill
	var employee entity.Employee
	var payment_types entity.PaymentType
	var foodordered entity.FoodOrdered
	var booking entity.Booking

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", bill.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10: ค้นหา Booking ด้วย id
	if tx := entity.DB().Where("id = ?", bill.BookingID).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	// 11: ค้นหา Foodered ด้วย id
	if tx := entity.DB().Where("id = ?", bill.FoodOrderedID).First(&foodordered); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodordered not found"})
		return
	}

	// 12: ค้นหา payment_types ด้วย id
	if tx := entity.DB().Where("id = ?", bill.PaymentTypeID).First(&payment_types); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentType not found"})
		return
	}

	// 13: สร้าง bill
	bill_pay := entity.Bill{
		Employee: employee,
		PaymentType: payment_types, // โยงความสัมพันธ์กับ Entity payment_types
		FoodOrdered: foodordered,   // โยงความสัมพันธ์กับ Entity place
		Booking:     booking,       // โยงความสัมพันธ์กับ Entity booking
		BillTime:    bill.BillTime, // ตั้งค่าฟิลด์ BillTime
		TotalPrice:  bill.TotalPrice, // ตั้งค่าฟิลด์ Total_price
	}

	// 14: บันทึก
	if err := entity.DB().Create(&bill_pay).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bill_pay})
}

// GET /bills
func ListBill(c *gin.Context) {
	var bill []entity.Bill
	if err := entity.DB().
		Preload("PaymentType").
		Preload("FoodOrdered").
		Preload("Booking").
		Preload("Booking.User").
		Preload("Employee").
		Raw("SELECT * FROM bills").Find(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// GET /Bill/:id  //ใข้ใน login

// func ListIDBill(c *gin.Context) {
// 	id := c.Param("id")
// 	var booking []entity.booking
// 	var bill []entity.Bill

// 	if tx := entity.DB().Where("user_id = ?", id).Find(&booking); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
// 		return
// 	}

// 	for _, reg := range booking {
// 		var b entity.Bill
// 		if tx := entity.DB().Preload("food_ordereds").
// 			Preload("PaymentType").
// 			Preload("Booking").
// 			Preload("Booking.User").
// 			Where("booking_id = ?", reg.ID).Find(&b); tx.RowsAffected == 0 {
// 			continue
// 		} else {
// 			bill = append(bill, b)
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": bill})
// }
