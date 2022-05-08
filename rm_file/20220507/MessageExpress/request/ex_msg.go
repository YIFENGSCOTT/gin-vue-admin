package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ExMsgSearch struct{
    MessageExpress.ExMsg
    request.PageInfo
}
