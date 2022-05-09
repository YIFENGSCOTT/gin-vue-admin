package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExMessageRouter struct {
}

// InitExMessageRouter 初始化 ExMessage 路由信息
func (s *ExMessageRouter) InitExMessageRouter(Router *gin.RouterGroup) {
	exMessageRouter := Router.Group("exMessage").Use(middleware.OperationRecord())
	exMessageRouterWithoutRecord := Router.Group("exMessage")
	var exMessageApi = v1.ApiGroupApp.MessageExpressApiGroup.ExMessageApi
	{
		exMessageRouter.POST("createExMessage", exMessageApi.CreateExMessage)             // 新建ExMessage
		exMessageRouter.DELETE("deleteExMessage", exMessageApi.DeleteExMessage)           // 删除ExMessage
		exMessageRouter.DELETE("deleteExMessageByIds", exMessageApi.DeleteExMessageByIds) // 批量删除ExMessage
		exMessageRouter.PUT("updateExMessage", exMessageApi.UpdateExMessage)              // 更新ExMessage
	}
	{
		exMessageRouterWithoutRecord.GET("findExMessage", exMessageApi.FindExMessage)             // 根据ID获取ExMessage
		exMessageRouterWithoutRecord.GET("findExMessageByCode", exMessageApi.FindExMessageByCode) // 根据code获取ExMessage
		exMessageRouterWithoutRecord.GET("findExMessageByPin", exMessageApi.FindExMessageByPin)   // 根据pin获取ExMessage
		exMessageRouterWithoutRecord.GET("getExMessageList", exMessageApi.GetExMessageList)       // 获取ExMessage列表
	}
}
