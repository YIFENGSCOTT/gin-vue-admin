// 自动生成模板ExMessage
package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExMessage 结构体
// 如果含有time.Time 请自行import time包
type ExMessage struct {
	global.GVA_MODEL
	Secret string `json:"secret" form:"secret" gorm:"column:secret;comment:;"`
	Pin    string `json:"pin" form:"pin" gorm:"column:pin;comment:;"`
	Code   string `json:"code" form:"code" gorm:"column:code;comment:;"`
	Saved  *bool  `json:"saved" form:"saved" gorm:"column:saved;comment:;"`
	Visits *int   `json:"visits" form:"visits" gorm:"column:visits;comment:;"`
	Xxx    string `json:"xxx" form:"xxx" gorm:"column:xxx;comment:;"`
}

// TableName ExMessage 表名
func (ExMessage) TableName() string {
	return "ex_message"
}
