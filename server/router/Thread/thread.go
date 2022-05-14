package Thread

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ThreadRouter struct {
}

// InitThreadRouter 初始化 Thread 路由信息
func (s *ThreadRouter) InitThreadRouter(Router *gin.RouterGroup) {
	threadRouter := Router.Group("thread").Use(middleware.OperationRecord())
	threadRouterWithoutRecord := Router.Group("thread")
	var threadApi = v1.ApiGroupApp.ThreadApiGroup.ThreadApi
	{
		threadRouter.POST("createThread", threadApi.CreateThread)             // 新建Thread
		threadRouter.DELETE("deleteThread", threadApi.DeleteThread)           // 删除Thread
		threadRouter.DELETE("deleteThreadByIds", threadApi.DeleteThreadByIds) // 批量删除Thread
		threadRouter.PUT("updateThread", threadApi.UpdateThread)              // 更新Thread
	}
	{
		threadRouterWithoutRecord.GET("findThread", threadApi.FindThread)         // 根据ID获取Thread
		threadRouterWithoutRecord.GET("findThreadByEK", threadApi.FindThreadByEK) // 根据ID获取Thread
		threadRouterWithoutRecord.GET("getThreadList", threadApi.GetThreadList)   // 获取Thread列表
	}
}
