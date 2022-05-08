// 自动生成模板ExMsg
package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExMsg 结构体
// 如果含有time.Time 请自行import time包
type ExMsg struct {
      global.GVA_MODEL
      Secret  string `json:"secret" form:"secret" gorm:"column:secret;comment:;"`
}


// TableName ExMsg 表名
func (ExMsg) TableName() string {
  return "exMessage"
}

