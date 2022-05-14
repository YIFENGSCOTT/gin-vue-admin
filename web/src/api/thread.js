import service from '@/utils/request'

// @Tags Thread
// @Summary 创建Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Thread true "创建Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thread/createThread [post]
export const createThread = (data) => {
  return service({
    url: '/thread/createThread',
    method: 'post',
    data
  })
}

// @Tags Thread
// @Summary 删除Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Thread true "删除Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /thread/deleteThread [delete]
export const deleteThread = (data) => {
  return service({
    url: '/thread/deleteThread',
    method: 'delete',
    data
  })
}

// @Tags Thread
// @Summary 删除Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /thread/deleteThread [delete]
export const deleteThreadByIds = (data) => {
  return service({
    url: '/thread/deleteThreadByIds',
    method: 'delete',
    data
  })
}

// @Tags Thread
// @Summary 更新Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Thread true "更新Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /thread/updateThread [put]
export const updateThread = (data) => {
  return service({
    url: '/thread/updateThread',
    method: 'put',
    data
  })
}

// @Tags Thread
// @Summary 用id查询Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Thread true "用id查询Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /thread/findThread [get]
export const findThread = (params) => {
  return service({
    url: '/thread/findThread',
    method: 'get',
    params
  })
}

// @Tags Thread
// @Summary 用EK查询Thread
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Thread true "用EK查询Thread"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /thread/findThread [get]
export const findThreadByEK = (params) => {
  return service({
    url: '/thread/findThreadByEK',
    method: 'get',
    params
  })
}

// @Tags Thread
// @Summary 分页获取Thread列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Thread列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thread/getThreadList [get]
export const getThreadList = (params) => {
  return service({
    url: '/thread/getThreadList',
    method: 'get',
    params
  })
}
