package entity

import (
	"bnqkl/chain-cms/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitEntityModule(db *gorm.DB, log *logger.Logger) {
	NewEntityService(db, log)
	NewEntityController()
}

func RegisterEntityApi(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("entity")
	// jwt 鉴权
	// group.Use(middlewares.JwtMiddleware())

	group.POST("/add", entityController.Add)
	group.POST("/add/multi", entityController.AddMulti)
	group.POST("/update", entityController.Update)
	group.GET("/factory/all", entityController.GetUserFactoryAll)
	group.GET("/factory/entity/all", entityController.GetUserFactoryEntityAll)
}
