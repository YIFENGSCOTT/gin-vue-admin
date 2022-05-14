package Thread

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Thread"
	ThreadReq "github.com/flipped-aurora/gin-vue-admin/server/model/Thread/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ThreadService struct {
}

// CreateThread 创建Thread记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) CreateThread(thread Thread.Thread) (err error) {
	err = global.GVA_DB.Create(&thread).Error
	return err
}

// DeleteThread 删除Thread记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) DeleteThread(thread Thread.Thread) (err error) {
	err = global.GVA_DB.Delete(&thread).Error
	return err
}

// DeleteThreadByIds 批量删除Thread记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) DeleteThreadByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Thread.Thread{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateThread 更新Thread记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) UpdateThread(thread Thread.Thread) (err error) {
	err = global.GVA_DB.Save(&thread).Error
	return err
}

// GetThread 根据id获取Thread记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) GetThread(id uint) (err error, thread Thread.Thread) {
	err = global.GVA_DB.Where("id = ?", id).First(&thread).Error
	return
}

// GetThreadByEK 根据EK获取Thread记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) GetThreadByEK(EK *int) (err error, threads []Thread.Thread) {
	err = global.GVA_DB.Where("encryption_key_id = ?", EK).Find(&threads).Error
	err = nil
	return
}

// GetThreadInfoList 分页获取Thread记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) GetThreadInfoList(info ThreadReq.ThreadSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Thread.Thread{})
	var threads []Thread.Thread
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ThreadContent != "" {
		db = db.Where("thread_content = ?", info.ThreadContent)
	}
	if info.FatherThreadId != nil {
		db = db.Where("father_thread = ?", info.FatherThreadId)
	}
	if info.EncryptionKeyId != nil {
		db = db.Where("encryption_key_id = ?", info.EncryptionKeyId)
	}
	if info.Beiyong != "" {
		db = db.Where("beiyong = ?", info.Beiyong)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&threads).Error
	return err, threads, total
}
