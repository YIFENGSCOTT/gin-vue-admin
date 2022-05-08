package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    MessageExpressReq "github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress/request"
)

type ExMsgService struct {
}

// CreateExMsg 创建ExMsg记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMsgService) CreateExMsg(exMessage MessageExpress.ExMsg) (err error) {
	err = global.GVA_DB.Create(&exMessage).Error
	return err
}

// DeleteExMsg 删除ExMsg记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMsgService)DeleteExMsg(exMessage MessageExpress.ExMsg) (err error) {
	err = global.GVA_DB.Delete(&exMessage).Error
	return err
}

// DeleteExMsgByIds 批量删除ExMsg记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMsgService)DeleteExMsgByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]MessageExpress.ExMsg{},"id in ?",ids.Ids).Error
	return err
}

// UpdateExMsg 更新ExMsg记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMsgService)UpdateExMsg(exMessage MessageExpress.ExMsg) (err error) {
	err = global.GVA_DB.Save(&exMessage).Error
	return err
}

// GetExMsg 根据id获取ExMsg记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMsgService)GetExMsg(id uint) (err error, exMessage MessageExpress.ExMsg) {
	err = global.GVA_DB.Where("id = ?", id).First(&exMessage).Error
	return
}

// GetExMsgInfoList 分页获取ExMsg记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMsgService)GetExMsgInfoList(info MessageExpressReq.ExMsgSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&MessageExpress.ExMsg{})
    var exMessages []MessageExpress.ExMsg
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&exMessages).Error
	return err, exMessages, total
}
