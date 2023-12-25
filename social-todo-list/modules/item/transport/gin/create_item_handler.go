package gin

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoaitemCreation
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
