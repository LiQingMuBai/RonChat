import service from '@/utils/request'
// @Tags RonOperatingActivity
// @Summary 创建ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonOperatingActivity true "创建ronOperatingActivity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ronOperatingActivity/createRonOperatingActivity [post]
export const createRonOperatingActivity = (data) => {
  return service({
    url: '/ronOperatingActivity/createRonOperatingActivity',
    method: 'post',
    data
  })
}

// @Tags RonOperatingActivity
// @Summary 删除ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonOperatingActivity true "删除ronOperatingActivity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ronOperatingActivity/deleteRonOperatingActivity [delete]
export const deleteRonOperatingActivity = (params) => {
  return service({
    url: '/ronOperatingActivity/deleteRonOperatingActivity',
    method: 'delete',
    params
  })
}

// @Tags RonOperatingActivity
// @Summary 批量删除ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ronOperatingActivity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ronOperatingActivity/deleteRonOperatingActivity [delete]
export const deleteRonOperatingActivityByIds = (params) => {
  return service({
    url: '/ronOperatingActivity/deleteRonOperatingActivityByIds',
    method: 'delete',
    params
  })
}

// @Tags RonOperatingActivity
// @Summary 更新ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonOperatingActivity true "更新ronOperatingActivity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ronOperatingActivity/updateRonOperatingActivity [put]
export const updateRonOperatingActivity = (data) => {
  return service({
    url: '/ronOperatingActivity/updateRonOperatingActivity',
    method: 'put',
    data
  })
}

// @Tags RonOperatingActivity
// @Summary 用id查询ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.RonOperatingActivity true "用id查询ronOperatingActivity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ronOperatingActivity/findRonOperatingActivity [get]
export const findRonOperatingActivity = (params) => {
  return service({
    url: '/ronOperatingActivity/findRonOperatingActivity',
    method: 'get',
    params
  })
}

// @Tags RonOperatingActivity
// @Summary 分页获取ronOperatingActivity表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ronOperatingActivity表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ronOperatingActivity/getRonOperatingActivityList [get]
export const getRonOperatingActivityList = (params) => {
  return service({
    url: '/ronOperatingActivity/getRonOperatingActivityList',
    method: 'get',
    params
  })
}

// @Tags RonOperatingActivity
// @Summary 不需要鉴权的ronOperatingActivity表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.RonOperatingActivitySearch true "分页获取ronOperatingActivity表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ronOperatingActivity/getRonOperatingActivityPublic [get]
export const getRonOperatingActivityPublic = () => {
  return service({
    url: '/ronOperatingActivity/getRonOperatingActivityPublic',
    method: 'get',
  })
}
