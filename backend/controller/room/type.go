package controller
import (
	"net/http"
	"github.com/Siriwan38/Sa-65-Group-18/entity"
	"github.com/gin-gonic/gin"
)

func ListType(c *gin.Context) {
	var Type []entity.Type
	if err := entity.DB().Table("types").Find(&Type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Type})
}
