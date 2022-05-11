package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Thread"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ThreadSearch struct {
	Thread.Thread
	request.PageInfo
}
