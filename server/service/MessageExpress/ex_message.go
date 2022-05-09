package MessageExpress

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress"
	MessageExpressReq "github.com/flipped-aurora/gin-vue-admin/server/model/MessageExpress/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ExMessageService struct {
}

// CreateExMessage 创建ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) CreateExMessage(exMessage MessageExpress.ExMessage) (err error) {
	err = global.GVA_DB.Create(&exMessage).Error
	return err
}

// DeleteExMessage 删除ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) DeleteExMessage(exMessage MessageExpress.ExMessage) (err error) {
	err = global.GVA_DB.Delete(&exMessage).Error
	return err
}

// DeleteExMessageByIds 批量删除ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) DeleteExMessageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]MessageExpress.ExMessage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExMessage 更新ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) UpdateExMessage(exMessage MessageExpress.ExMessage) (err error) {
	err = global.GVA_DB.Save(&exMessage).Error
	return err
}

// GetExMessage 根据id获取ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) GetExMessage(id uint) (err error, exMessage MessageExpress.ExMessage) {
	err = global.GVA_DB.Where("id = ?", id).First(&exMessage).Error
	return
}

// GetExMessage 根据code获取ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) GetExMessageByCode(code string) (err error, exMessage MessageExpress.ExMessage) {
	err = global.GVA_DB.Where("code = ?", code).First(&exMessage).Error
	return
}

// GetExMessage 根据pin获取ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) GetExMessageByPin(pin string) (err error, exMessage MessageExpress.ExMessage) {
	err = global.GVA_DB.Where("pin = ?", pin).First(&exMessage).Error
	return
}

// GetExMessageInfoList 分页获取ExMessage记录
// Author [piexlmax](https://github.com/piexlmax)
func (exMessageService *ExMessageService) GetExMessageInfoList(info MessageExpressReq.ExMessageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&MessageExpress.ExMessage{})
	var exMessages []MessageExpress.ExMessage
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&exMessages).Error
	return err, exMessages, total
}
