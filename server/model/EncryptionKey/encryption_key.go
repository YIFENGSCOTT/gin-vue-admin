// 自动生成模板EncryptionKey
package EncryptionKey

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EncryptionKey 结构体
// 如果含有time.Time 请自行import time包
type EncryptionKey struct {
	global.GVA_MODEL
	UserId             *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
	KeyContent         string `json:"keyContent" form:"keyContent" gorm:"column:key_content;comment:;"`
	KeyIllustrationUrl string `json:"keyIllustrationUrl" form:"keyIllustrationUrl" gorm:"column:key_illustration_url;comment:;"`
	Beiyong            string `json:"beiyong" form:"beiyong" gorm:"column:beiyong;comment:;"`
}

// TableName EncryptionKey 表名
func (EncryptionKey) TableName() string {
	return "encryption_key"
}
