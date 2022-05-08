import service from '@/utils/request'

// @Tags ExMsg
// @Summary 创建ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExMsg true "创建ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/createExMsg [post]
export const createExMsg = (data) => {
  return service({
    url: '/exMessage/createExMsg',
    method: 'post',
    data
  })
}

// @Tags ExMsg
// @Summary 删除ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExMsg true "删除ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exMessage/deleteExMsg [delete]
export const deleteExMsg = (data) => {
  return service({
    url: '/exMessage/deleteExMsg',
    method: 'delete',
    data
  })
}

// @Tags ExMsg
// @Summary 删除ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exMessage/deleteExMsg [delete]
export const deleteExMsgByIds = (data) => {
  return service({
    url: '/exMessage/deleteExMsgByIds',
    method: 'delete',
    data
  })
}

// @Tags ExMsg
// @Summary 更新ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExMsg true "更新ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exMessage/updateExMsg [put]
export const updateExMsg = (data) => {
  return service({
    url: '/exMessage/updateExMsg',
    method: 'put',
    data
  })
}

// @Tags ExMsg
// @Summary 用id查询ExMsg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExMsg true "用id查询ExMsg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMsg [get]
export const findExMsg = (params) => {
  return service({
    url: '/exMessage/findExMsg',
    method: 'get',
    params
  })
}

// @Tags ExMsg
// @Summary 分页获取ExMsg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExMsg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/getExMsgList [get]
export const getExMsgList = (params) => {
  return service({
    url: '/exMessage/getExMsgList',
    method: 'get',
    params
  })
}
