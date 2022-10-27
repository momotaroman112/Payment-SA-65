package controller

import (
	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)
// POST /equipment
func CreateEquipment(c *gin.Context) {
	var category entity.Category
	var employee entity.Employee
	var unit entity.Unit
	var equipment entity.Equipment
	//ผลลัพทธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปล equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//9: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", equipment.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}
	//10: ค้นหา Category ด้วย id
	if tx := entity.DB().Where("id = ?", equipment.CategoryID).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}
	//11: ค้นหา Unit ด้วย id
	if tx := entity.DB().Where("id = ?", equipment.UnitID).First(&unit); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unit not found"})
		return
	}
	//12: สร้าง Equipment
	eq := entity.Equipment{
		Category: category,
		Unit:     unit,
		Employee: employee,
		Time:     equipment.Time,
		Name:     equipment.Name,
		Amount:   equipment.Amount,
	}
	//13: บันทึก
	if err := entity.DB().Create(&eq).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": eq})
}
// GET /equipment/:id
func GetEquipment(c *gin.Context) {
	var equipment entity.Equipment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM equipment WHERE id = ?", id).
		Preload("Category").
		Preload("Unit").
		Preload("Employee").
		Find(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": equipment})
}
// GET /equipments
func ListEquipments(c *gin.Context) {
	var equipments []entity.Equipment
	if err := entity.DB().Raw("SELECT * FROM equipment").
		Preload("Category").
		Preload("Unit").
		Preload("Employee").
		Find(&equipments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": equipments})
}
// DELETE /equipment/:id
func DeleteEquipment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM equipment WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
// PATCH /equipment
func UpdateEquipment(c *gin.Context) {
	var equipment entity.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", equipment.ID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}
	if err := entity.DB().Save(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": equipment})
}