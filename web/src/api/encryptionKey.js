import service from '@/utils/request'

// @Tags EncryptionKey
// @Summary 创建EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptionKey true "创建EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /encryptionKey/createEncryptionKey [post]
export const createEncryptionKey = (data) => {
  return service({
    url: '/encryptionKey/createEncryptionKey',
    method: 'post',
    data
  })
}

// @Tags EncryptionKey
// @Summary 删除EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptionKey true "删除EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /encryptionKey/deleteEncryptionKey [delete]
export const deleteEncryptionKey = (data) => {
  return service({
    url: '/encryptionKey/deleteEncryptionKey',
    method: 'delete',
    data
  })
}

// @Tags EncryptionKey
// @Summary 删除EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /encryptionKey/deleteEncryptionKey [delete]
export const deleteEncryptionKeyByIds = (data) => {
  return service({
    url: '/encryptionKey/deleteEncryptionKeyByIds',
    method: 'delete',
    data
  })
}

// @Tags EncryptionKey
// @Summary 更新EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptionKey true "更新EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /encryptionKey/updateEncryptionKey [put]
export const updateEncryptionKey = (data) => {
  return service({
    url: '/encryptionKey/updateEncryptionKey',
    method: 'put',
    data
  })
}

// @Tags EncryptionKey
// @Summary 用id查询EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EncryptionKey true "用id查询EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /encryptionKey/findEncryptionKey [get]
export const findEncryptionKey = (params) => {
  return service({
    url: '/encryptionKey/findEncryptionKey',
    method: 'get',
    params
  })
}

// @Tags EncryptionKeyByContent
// @Summary 用content查询EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EncryptionKey true "用key_content查询EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /encryptionKey/findEncryptionKey [get]
export const findEncryptionKeyByContent = (params) => {
  return service({
    url: '/encryptionKey/findEncryptionKeyByContent',
    method: 'get',
    params
  })
}

// @Tags EncryptionKey
// @Summary 分页获取EncryptionKey列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EncryptionKey列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /encryptionKey/getEncryptionKeyList [get]
export const getEncryptionKeyList = (params) => {
  return service({
    url: '/encryptionKey/getEncryptionKeyList',
    method: 'get',
    params
  })
}
