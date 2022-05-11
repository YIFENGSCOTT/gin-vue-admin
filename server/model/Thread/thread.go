// 自动生成模板Thread
package Thread

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Thread 结构体
// 如果含有time.Time 请自行import time包
type Thread struct {
	global.GVA_MODEL
	ThreadContent   string `json:"threadContent" form:"threadContent" gorm:"column:thread_content;comment:;"`
	FatherThreadId  *int   `json:"fatherThread" form:"fatherThread" gorm:"column:father_thread;comment:;"`
	EncryptionKeyId *int   `json:"encryptionKeyId" form:"encryptionKeyId" gorm:"column:encryption_key_id;comment:;"`
	Visited         *int   `json:"visited" form:"visited" gorm:"column:visited;comment:;"`
	Beiyong         string `json:"beiyong" form:"beiyong" gorm:"column:beiyong;comment:;"`
}

// TableName Thread 表名
func (Thread) TableName() string {
	return "Thread"
}
