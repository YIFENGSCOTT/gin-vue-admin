package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/EncryptionKey"
	"github.com/flipped-aurora/gin-vue-admin/server/router/MessageExpress"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Thread"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System         system.RouterGroup
	Example        example.RouterGroup
	MessageExpress MessageExpress.RouterGroup
	EncryptionKey  EncryptionKey.RouterGroup
	Thread         Thread.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
