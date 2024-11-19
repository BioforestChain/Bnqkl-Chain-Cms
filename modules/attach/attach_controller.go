package attach

import (
	"bnqkl/chain-cms/helper"

	"github.com/gin-gonic/gin"
)

type AttachController struct{}

var attachController *AttachController

func NewAttachController() {
	attachController = &AttachController{}
}

func GetAttachController() *AttachController {
	return attachController
}

// @Tags 附件模块
// @Summary 上传文件
// @Accept application/json
// @Produce application/json
// @Param name formData string true "blob名称"
// @Param extension formData string false "blob扩展名"
// @Param file formData file true "blob内容"
// @Success 200	{object} UploadBlobRes
// @Router /attach/upload/blob [post]
func (c *AttachController) UploadBlob(ctx *gin.Context) {
	var req UploadBlobReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		helper.FailureResponse(ctx, err)
		return
	}
	resp, err := attachService.UploadBlob(req)
	helper.ResponseWrapper(ctx, resp, err)
}
