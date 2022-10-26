package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /foodsets
func CreateFoodSet(c *gin.Context) {
	var foodset entity.FoodSet
	if err := c.ShouldBindJSON(&foodset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&foodset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foodset})
}

// GET /foodset/:id
func GetFoodSet(c *gin.Context) {
	var foodset entity.FoodSet

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM food_sets WHERE id = ?", id).Find(&foodset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodset})
}

// GET /foodsets
func ListFoodSets(c *gin.Context) {
	var foodsets []entity.FoodSet //[] ส่งเป็นแบบลิสต์

	if err := entity.DB().Raw("SELECT * FROM food_sets").Scan(&foodsets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodsets})
}

// DELETE /foodsets/:id
// func DeleteFoodSet(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM food_sets WHERE id = ?", id); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "food set not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }

// PATCH /foodsets
// func UpdateFoodSet(c *gin.Context) {
// 	var foodset entity.FoodSet
// 	if err := c.ShouldBindJSON(&foodset); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entity.DB().Where("id = ?", foodset.ID).First(&foodset); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "food set not found"})
// 		return
// 	}

// 	if err := entity.DB().Save(&foodset).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": foodset})
// }

/* Note.
1. FoodSet เป็นตารางรอง ไม่จำเป็นต้องมี preload เพราะไม่ต้องไปดึงของใครมาใส่ของตัวเอง
2.
*/
