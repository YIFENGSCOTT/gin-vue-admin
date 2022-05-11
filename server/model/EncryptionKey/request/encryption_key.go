package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/EncryptionKey"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EncryptionKeySearch struct {
	EncryptionKey.EncryptionKey
	request.PageInfo
}
