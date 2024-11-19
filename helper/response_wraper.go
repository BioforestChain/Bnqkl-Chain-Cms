package helper

import (
	"bnqkl/chain-cms/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  data,
	})
}

func FailureResponse(ctx *gin.Context, err error) {
	rErr, ok := err.(*exception.ResponseError)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   rErr,
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   *exception.NewErrorCode("500", err.Error()),
	})
}

func ResponseWrapper[T any](ctx *gin.Context, data T, err error) {
	if err != nil {
		FailureResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, data)
}
