package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/EncryptionKey"
	"github.com/flipped-aurora/gin-vue-admin/server/service/MessageExpress"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Thread"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	MessageExpressServiceGroup MessageExpress.ServiceGroup
	EncryptionKeyServiceGroup  EncryptionKey.ServiceGroup
	ThreadServiceGroup         Thread.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
