package entity

import (
	"bnqkl/chain-cms/helper"

	"github.com/gin-gonic/gin"
)

type EntityController struct{}

var entityController *EntityController

func NewEntityController() {
	entityController = &EntityController{}
}

func GetEntityController() *EntityController {
	return entityController
}

// @Tags entity 模块
// @Summary 新增 entity
// @Accept application/json
// @Produce pplication/json
// @Param object body AddEntityReq true "请求参数"
// @Success 200	{object} AddEntityRes
// @Router /entity/add [post]
func (mc *EntityController) Add(ctx *gin.Context) {
	var req AddEntityReq
	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		helper.FailureResponse(ctx, err)
		return
	}
	resp, err := entityService.Add(req)
	helper.ResponseWrapper(ctx, resp, err)
}

// @Tags entity 模块
// @Summary 批量新增 entity
// @Accept application/json
// @Produce pplication/json
// @Param object body AddEntityMultiReq true "请求参数"
// @Success 200	{object} AddEntityMultiRes
// @Router /entity/add/multi [post]
func (mc *EntityController) AddMulti(ctx *gin.Context) {
	var req AddEntityMultiReq
	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		helper.FailureResponse(ctx, err)
		return
	}
	resp, err := entityService.AddMulti(req)
	helper.ResponseWrapper(ctx, resp, err)
}

// @Tags entity 模块
// @Summary 更新 entity
// @Accept application/json
// @Produce pplication/json
// @Param object body UpdateEntityReq true "请求参数"
// @Success 200	{object} UpdateEntityRes
// @Router /entity/update [post]
func (mc *EntityController) Update(ctx *gin.Context) {
	var req UpdateEntityReq
	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		helper.FailureResponse(ctx, err)
		return
	}
	resp, err := entityService.Update(req)
	helper.ResponseWrapper(ctx, resp, err)
}

// @Tags entity 模块
// @Summary 获取用户持有的 factory 总览
// @Accept application/json
// @Produce application/json
// @Param object query GetUserFactoryAllReq true "请求参数"
// @Success 200	{object} GetUserFactoryAllRes
// @Router /entity/factory/all [get]
func (mc *EntityController) GetUserFactoryAll(ctx *gin.Context) {
	var req GetUserFactoryAllReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		helper.FailureResponse(ctx, err)
		return
	}
	resp, err := entityService.GetUserFactoryAll(req)
	helper.ResponseWrapper(ctx, resp, err)
}

// @Tags entity 模块
// @Summary 获取用户持有的(某个 factory 下的)所有 entity
// @Accept application/json
// @Produce application/json
// @Param object query GetUserFactoryEntityAllReq true "请求参数"
// @Success 200	{object} GetUserFactoryEntityAllRes
// @Router /entity/factory/entity/all [get]
func (mc *EntityController) GetUserFactoryEntityAll(ctx *gin.Context) {
	var req GetUserFactoryEntityAllReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		helper.FailureResponse(ctx, err)
		return
	}
	resp, err := entityService.GetUserFactoryEntityAll(req)
	helper.ResponseWrapper(ctx, resp, err)
}
