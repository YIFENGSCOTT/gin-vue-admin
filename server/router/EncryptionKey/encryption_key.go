package EncryptionKey

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EncryptionKeyRouter struct {
}

// InitEncryptionKeyRouter 初始化 EncryptionKey 路由信息
func (s *EncryptionKeyRouter) InitEncryptionKeyRouter(Router *gin.RouterGroup) {
	encryptionKeyRouter := Router.Group("encryptionKey").Use(middleware.OperationRecord())
	encryptionKeyRouterWithoutRecord := Router.Group("encryptionKey")
	var encryptionKeyApi = v1.ApiGroupApp.EncryptionKeyApiGroup.EncryptionKeyApi
	{
		encryptionKeyRouter.POST("createEncryptionKey", encryptionKeyApi.CreateEncryptionKey)             // 新建EncryptionKey
		encryptionKeyRouter.DELETE("deleteEncryptionKey", encryptionKeyApi.DeleteEncryptionKey)           // 删除EncryptionKey
		encryptionKeyRouter.DELETE("deleteEncryptionKeyByIds", encryptionKeyApi.DeleteEncryptionKeyByIds) // 批量删除EncryptionKey
		encryptionKeyRouter.PUT("updateEncryptionKey", encryptionKeyApi.UpdateEncryptionKey)              // 更新EncryptionKey
	}
	{
		encryptionKeyRouterWithoutRecord.GET("findEncryptionKey", encryptionKeyApi.FindEncryptionKey)       // 根据ID获取EncryptionKey
		encryptionKeyRouterWithoutRecord.GET("getEncryptionKeyList", encryptionKeyApi.GetEncryptionKeyList) // 获取EncryptionKey列表
	}
}
