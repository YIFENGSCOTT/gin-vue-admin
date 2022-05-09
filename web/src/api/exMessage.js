import service from '@/utils/request'

// @Tags ExMessage
// @Summary 创建ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExMessage true "创建ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/createExMessage [post]
export const createExMessage = (data) => {
  return service({
    url: '/exMessage/createExMessage',
    method: 'post',
    data
  })
}

// @Tags ExMessage
// @Summary 删除ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExMessage true "删除ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exMessage/deleteExMessage [delete]
export const deleteExMessage = (data) => {
  return service({
    url: '/exMessage/deleteExMessage',
    method: 'delete',
    data
  })
}

// @Tags ExMessage
// @Summary 删除ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exMessage/deleteExMessage [delete]
export const deleteExMessageByIds = (data) => {
  return service({
    url: '/exMessage/deleteExMessageByIds',
    method: 'delete',
    data
  })
}

// @Tags ExMessage
// @Summary 更新ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExMessage true "更新ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exMessage/updateExMessage [put]
export const updateExMessage = (data) => {
  return service({
    url: '/exMessage/updateExMessage',
    method: 'put',
    data
  })
}

// @Tags ExMessage
// @Summary 用id查询ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExMessage true "用id查询ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMessage [get]
export const findExMessage = (params) => {
  return service({
    url: '/exMessage/findExMessage',
    method: 'get',
    params
  })
}
// @Tags ExMessage
// @Summary 用code查询ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExMessage true "用code查询ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMessageByCode [get]
export const findExMessageByCode = (params) => {
  return service({
    url: '/exMessage/findExMessageByCode',
    method: 'get',
    params
  })
}
// @Tags ExMessage
// @Summary 用pin查询ExMessage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExMessage true "用pin查询ExMessage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exMessage/findExMessageByPin [get]
export const findExMessageByPin = (params) => {
  return service({
    url: '/exMessage/findExMessageByPin',
    method: 'get',
    params
  })
}

// @Tags ExMessage
// @Summary 分页获取ExMessage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExMessage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exMessage/getExMessageList [get]
export const getExMessageList = (params) => {
  return service({
    url: '/exMessage/getExMessageList',
    method: 'get',
    params
  })
}
