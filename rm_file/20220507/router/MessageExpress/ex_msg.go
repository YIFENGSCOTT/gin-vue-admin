package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExMsgRouter struct {
}

// InitExMsgRouter 初始化 ExMsg 路由信息
func (s *ExMsgRouter) InitExMsgRouter(Router *gin.RouterGroup) {
	exMessageRouter := Router.Group("exMessage").Use(middleware.OperationRecord())
	exMessageRouterWithoutRecord := Router.Group("exMessage")
	var exMessageApi = v1.ApiGroupApp.MessageExpressApiGroup.ExMsgApi
	{
		exMessageRouter.POST("createExMsg", exMessageApi.CreateExMsg)   // 新建ExMsg
		exMessageRouter.DELETE("deleteExMsg", exMessageApi.DeleteExMsg) // 删除ExMsg
		exMessageRouter.DELETE("deleteExMsgByIds", exMessageApi.DeleteExMsgByIds) // 批量删除ExMsg
		exMessageRouter.PUT("updateExMsg", exMessageApi.UpdateExMsg)    // 更新ExMsg
	}
	{
		exMessageRouterWithoutRecord.GET("findExMsg", exMessageApi.FindExMsg)        // 根据ID获取ExMsg
		exMessageRouterWithoutRecord.GET("getExMsgList", exMessageApi.GetExMsgList)  // 获取ExMsg列表
	}
}
