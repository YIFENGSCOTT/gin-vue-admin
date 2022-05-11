package EncryptionKey

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EncryptionKey"
	EncryptionKeyReq "github.com/flipped-aurora/gin-vue-admin/server/model/EncryptionKey/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EncryptionKeyService struct {
}

// CreateEncryptionKey 创建EncryptionKey记录
// Author [piexlmax](https://github.com/piexlmax)
func (encryptionKeyService *EncryptionKeyService) CreateEncryptionKey(encryptionKey EncryptionKey.EncryptionKey) (err error) {
	err = global.GVA_DB.Create(&encryptionKey).Error
	return err
}

// DeleteEncryptionKey 删除EncryptionKey记录
// Author [piexlmax](https://github.com/piexlmax)
func (encryptionKeyService *EncryptionKeyService) DeleteEncryptionKey(encryptionKey EncryptionKey.EncryptionKey) (err error) {
	err = global.GVA_DB.Delete(&encryptionKey).Error
	return err
}

// DeleteEncryptionKeyByIds 批量删除EncryptionKey记录
// Author [piexlmax](https://github.com/piexlmax)
func (encryptionKeyService *EncryptionKeyService) DeleteEncryptionKeyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EncryptionKey.EncryptionKey{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEncryptionKey 更新EncryptionKey记录
// Author [piexlmax](https://github.com/piexlmax)
func (encryptionKeyService *EncryptionKeyService) UpdateEncryptionKey(encryptionKey EncryptionKey.EncryptionKey) (err error) {
	err = global.GVA_DB.Save(&encryptionKey).Error
	return err
}

// GetEncryptionKey 根据id获取EncryptionKey记录
// Author [piexlmax](https://github.com/piexlmax)
func (encryptionKeyService *EncryptionKeyService) GetEncryptionKey(id uint) (err error, encryptionKey EncryptionKey.EncryptionKey) {
	err = global.GVA_DB.Where("id = ?", id).First(&encryptionKey).Error
	return
}

// GetEncryptionKeyInfoList 分页获取EncryptionKey记录
// Author [piexlmax](https://github.com/piexlmax)
func (encryptionKeyService *EncryptionKeyService) GetEncryptionKeyInfoList(info EncryptionKeyReq.EncryptionKeySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&EncryptionKey.EncryptionKey{})
	var encryptionKeys []EncryptionKey.EncryptionKey
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.KeyContent != "" {
		db = db.Where("key_content = ?", info.KeyContent)
	}
	if info.KeyIllustrationUrl != "" {
		db = db.Where("key_illustration_url = ?", info.KeyIllustrationUrl)
	}
	if info.Beiyong != "" {
		db = db.Where("beiyong = ?", info.Beiyong)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&encryptionKeys).Error
	return err, encryptionKeys, total
}
