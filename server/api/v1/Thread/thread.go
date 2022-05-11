package Thread

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Thread"
	ThreadReq "github.com/flipped-aurora/gin-vue-admin/server/model/Thread/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ThreadApi struct {
}

var threadService = service.ServiceGroupApp.ThreadServiceGroup.ThreadService

// CreateThread 创建Thread
// @Tags Thread
// @Summary 创建Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Thread.Thread true "创建Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thread/createThread [post]
func (threadApi *ThreadApi) CreateThread(c *gin.Context) {
	var thread Thread.Thread
	_ = c.ShouldBindJSON(&thread)
	if err := threadService.CreateThread(thread); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteThread 删除Thread
// @Tags Thread
// @Summary 删除Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Thread.Thread true "删除Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /thread/deleteThread [delete]
func (threadApi *ThreadApi) DeleteThread(c *gin.Context) {
	var thread Thread.Thread
	_ = c.ShouldBindJSON(&thread)
	if err := threadService.DeleteThread(thread); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteThreadByIds 批量删除Thread
// @Tags Thread
// @Summary 批量删除Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /thread/deleteThreadByIds [delete]
func (threadApi *ThreadApi) DeleteThreadByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := threadService.DeleteThreadByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateThread 更新Thread
// @Tags Thread
// @Summary 更新Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Thread.Thread true "更新Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /thread/updateThread [put]
func (threadApi *ThreadApi) UpdateThread(c *gin.Context) {
	var thread Thread.Thread
	_ = c.ShouldBindJSON(&thread)
	if err := threadService.UpdateThread(thread); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindThread 用id查询Thread
// @Tags Thread
// @Summary 用id查询Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Thread.Thread true "用id查询Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /thread/findThread [get]
func (threadApi *ThreadApi) FindThread(c *gin.Context) {
	var thread Thread.Thread
	_ = c.ShouldBindQuery(&thread)
	if err, rethread := threadService.GetThread(thread.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rethread": rethread}, c)
	}
}

// GetThreadList 分页获取Thread列表
// @Tags Thread
// @Summary 分页获取Thread列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ThreadReq.ThreadSearch true "分页获取Thread列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thread/getThreadList [get]
func (threadApi *ThreadApi) GetThreadList(c *gin.Context) {
	var pageInfo ThreadReq.ThreadSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := threadService.GetThreadInfoList(pageInfo); err != nil {
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
