package controller

import (
	"net/http"

	"github.com/AKiRA3563/sa-65/entity"
	"github.com/gin-gonic/gin"
)

// POST /createborrowlist
func CreateBorrowList(c *gin.Context) {

	var User entity.User
	var Equipment entity.Equipment
	var Employee entity.Employee
	var BorrowList entity.BorrowList

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร BorrowList
	if err := c.ShouldBindJSON(&BorrowList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//10: ค้นหา User ด้วย id
	if tx := entity.DB().Table("Users").Where("id = ?", BorrowList.MemberID).First(&User); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 11: ค้นหา Equipment ด้วย id
	if tx := entity.DB().Table("equipment").Where("id = ?", BorrowList.EquipmentID).First(&Equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	// 12: ค้นหา Employee ด้วย id
	if tx := entity.DB().Table("employees").Where("id = ?", BorrowList.EmployeeID).First(&Employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 14: สร้าง BorrowList
	lr := entity.BorrowList{
		Equipment:  Equipment,
		Member:		User,
		Amount:     BorrowList.Amount,
		Employee:   Employee,
		BorrowTime: BorrowList.BorrowTime,
	}

	// 15: บันทึก
	if err := entity.DB().Create(&lr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lr})
}

// GET: /borrowlist
func ListBorrow(c *gin.Context) {
	var BorrowList []*entity.BorrowList
	if err := entity.DB().Table("borrow_lists").Preload("Equipment").Preload("Employee").Preload("Member").Find(&BorrowList).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": BorrowList})
}
