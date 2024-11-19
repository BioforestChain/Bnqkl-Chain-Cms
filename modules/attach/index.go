package attach

import (
	"bnqkl/chain-cms/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAttachModule(db *gorm.DB, log *logger.Logger) {
	NewAttachService(db, log)
	NewAttachController()
}

func RegisterAttachApi(parentGroup *gin.RouterGroup) {
	group := parentGroup.Group("attach")
	// jwt 鉴权
	// group.Use(middlewares.JwtMiddleware())

	group.POST("/upload/blob", attachController.UploadBlob)
}
