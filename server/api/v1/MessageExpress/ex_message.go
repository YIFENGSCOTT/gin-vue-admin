package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress"
	MessageExpressReq "github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExMessageApi struct {
}

var exMessageService = service.ServiceGroupApp.MessageExpressServiceGroup.ExMessageService

// CreateExMessage 创建ExMessage
// @Tags ExMessage
// @Summary 创建ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MessageExpress.ExMessage true "创建ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/createExMessage [post]
func (exMessageApi *ExMessageApi) CreateExMessage(c *gin.Context) {
	var exMessage MessageExpress.ExMessage
	_ = c.ShouldBindJSON(&exMessage)
	// todo：加一段逻辑来生成code和pin
	exMessage.Code = utils.GenerateCode()
	exMessage.Pin = utils.BcryptHash(exMessage.Secret)
	int0 := 0
	exMessage.Visits = &int0
	if err := exMessageService.CreateExMessage(exMessage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(gin.H{"pin": exMessage.Pin, "code": exMessage.Code}, c)
	}
}

// DeleteExMessage 删除ExMessage
// @Tags ExMessage
// @Summary 删除ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MessageExpress.ExMessage true "删除ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exMessage/deleteExMessage [delete]
func (exMessageApi *ExMessageApi) DeleteExMessage(c *gin.Context) {
	var exMessage MessageExpress.ExMessage
	_ = c.ShouldBindJSON(&exMessage)
	if err := exMessageService.DeleteExMessage(exMessage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExMessageByIds 批量删除ExMessage
// @Tags ExMessage
// @Summary 批量删除ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exMessage/deleteExMessageByIds [delete]
func (exMessageApi *ExMessageApi) DeleteExMessageByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := exMessageService.DeleteExMessageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExMessage 更新ExMessage
// @Tags ExMessage
// @Summary 更新ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MessageExpress.ExMessage true "更新ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exMessage/updateExMessage [put]
func (exMessageApi *ExMessageApi) UpdateExMessage(c *gin.Context) {
	var exMessage MessageExpress.ExMessage
	_ = c.ShouldBindJSON(&exMessage)
	if err := exMessageService.UpdateExMessage(exMessage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExMessage 用id查询ExMessage
// @Tags ExMessage
// @Summary 用id查询ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MessageExpress.ExMessage true "用id查询ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMessage [get]
func (exMessageApi *ExMessageApi) FindExMessage(c *gin.Context) {
	var exMessage MessageExpress.ExMessage
	_ = c.ShouldBindQuery(&exMessage)
	if err, reexMessage := exMessageService.GetExMessage(exMessage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexMessage": reexMessage}, c)
	}
}

// FindExMessageByCode 用code查询ExMessage
// @Tags ExMessage
// @Summary 用code查询ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MessageExpress.ExMessage true "用code查询ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMessageByCode [get]
func (exMessageApi *ExMessageApi) FindExMessageByCode(c *gin.Context) {
	var exMessage MessageExpress.ExMessage
	_ = c.ShouldBindQuery(&exMessage)
	if err, reexMessage := exMessageService.GetExMessageByCode(exMessage.Code); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexMessage": reexMessage}, c)
	}
}

// FindExMessageByPin 用pin查询ExMessage
// @Tags ExMessage
// @Summary 用pin查询ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MessageExpress.ExMessage true "用pin查询ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMessageByPin [get]
func (exMessageApi *ExMessageApi) FindExMessageByPin(c *gin.Context) {
	var exMessage MessageExpress.ExMessage
	_ = c.ShouldBindQuery(&exMessage)
	if err, reexMessage := exMessageService.GetExMessageByPin(exMessage.Pin); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexMessage": reexMessage}, c)
	}
}

// GetExMessageList 分页获取ExMessage列表
// @Tags ExMessage
// @Summary 分页获取ExMessage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MessageExpressReq.ExMessageSearch true "分页获取ExMessage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/getExMessageList [get]
func (exMessageApi *ExMessageApi) GetExMessageList(c *gin.Context) {
	var pageInfo MessageExpressReq.ExMessageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := exMessageService.GetExMessageInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
