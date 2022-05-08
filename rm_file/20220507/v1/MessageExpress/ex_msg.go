package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    MessageExpressReq "github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ExMsgApi struct {
}

var exMessageService = service.ServiceGroupApp.MessageExpressServiceGroup.ExMsgService


// CreateExMsg 创建ExMsg
// @Tags ExMsg
// @Summary 创建ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MessageExpress.ExMsg true "创建ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/createExMsg [post]
func (exMessageApi *ExMsgApi) CreateExMsg(c *gin.Context) {
	var exMessage MessageExpress.ExMsg
	_ = c.ShouldBindJSON(&exMessage)
	if err := exMessageService.CreateExMsg(exMessage); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExMsg 删除ExMsg
// @Tags ExMsg
// @Summary 删除ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MessageExpress.ExMsg true "删除ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exMessage/deleteExMsg [delete]
func (exMessageApi *ExMsgApi) DeleteExMsg(c *gin.Context) {
	var exMessage MessageExpress.ExMsg
	_ = c.ShouldBindJSON(&exMessage)
	if err := exMessageService.DeleteExMsg(exMessage); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExMsgByIds 批量删除ExMsg
// @Tags ExMsg
// @Summary 批量删除ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exMessage/deleteExMsgByIds [delete]
func (exMessageApi *ExMsgApi) DeleteExMsgByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := exMessageService.DeleteExMsgByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExMsg 更新ExMsg
// @Tags ExMsg
// @Summary 更新ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MessageExpress.ExMsg true "更新ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exMessage/updateExMsg [put]
func (exMessageApi *ExMsgApi) UpdateExMsg(c *gin.Context) {
	var exMessage MessageExpress.ExMsg
	_ = c.ShouldBindJSON(&exMessage)
	if err := exMessageService.UpdateExMsg(exMessage); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExMsg 用id查询ExMsg
// @Tags ExMsg
// @Summary 用id查询ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MessageExpress.ExMsg true "用id查询ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMsg [get]
func (exMessageApi *ExMsgApi) FindExMsg(c *gin.Context) {
	var exMessage MessageExpress.ExMsg
	_ = c.ShouldBindQuery(&exMessage)
	if err, reexMessage := exMessageService.GetExMsg(exMessage.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexMessage": reexMessage}, c)
	}
}

// GetExMsgList 分页获取ExMsg列表
// @Tags ExMsg
// @Summary 分页获取ExMsg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MessageExpressReq.ExMsgSearch true "分页获取ExMsg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/getExMsgList [get]
func (exMessageApi *ExMsgApi) GetExMsgList(c *gin.Context) {
	var pageInfo MessageExpressReq.ExMsgSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := exMessageService.GetExMsgInfoList(pageInfo); err != nil {
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
