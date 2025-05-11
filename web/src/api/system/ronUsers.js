import service from '@/utils/request'
// @Tags RonUsers
// @Summary 创建ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonUsers true "创建ronUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ronUsers/createRonUsers [post]
export const createRonUsers = (data) => {
  return service({
    url: '/ronUsers/createRonUsers',
    method: 'post',
    data
  })
}

// @Tags RonUsers
// @Summary 删除ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonUsers true "删除ronUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ronUsers/deleteRonUsers [delete]
export const deleteRonUsers = (params) => {
  return service({
    url: '/ronUsers/deleteRonUsers',
    method: 'delete',
    params
  })
}

// @Tags RonUsers
// @Summary 批量删除ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ronUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ronUsers/deleteRonUsers [delete]
export const deleteRonUsersByIds = (params) => {
  return service({
    url: '/ronUsers/deleteRonUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags RonUsers
// @Summary 更新ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonUsers true "更新ronUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ronUsers/updateRonUsers [put]
export const updateRonUsers = (data) => {
  return service({
    url: '/ronUsers/updateRonUsers',
    method: 'put',
    data
  })
}

// @Tags RonUsers
// @Summary 用id查询ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.RonUsers true "用id查询ronUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ronUsers/findRonUsers [get]
export const findRonUsers = (params) => {
  return service({
    url: '/ronUsers/findRonUsers',
    method: 'get',
    params
  })
}

// @Tags RonUsers
// @Summary 分页获取ronUsers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ronUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ronUsers/getRonUsersList [get]
export const getRonUsersList = (params) => {
  return service({
    url: '/ronUsers/getRonUsersList',
    method: 'get',
    params
  })
}

// @Tags RonUsers
// @Summary 不需要鉴权的ronUsers表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.RonUsersSearch true "分页获取ronUsers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ronUsers/getRonUsersPublic [get]
export const getRonUsersPublic = () => {
  return service({
    url: '/ronUsers/getRonUsersPublic',
    method: 'get',
  })
}
