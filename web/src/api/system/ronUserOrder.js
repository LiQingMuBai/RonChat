import service from '@/utils/request'
// @Tags RonUserOrder
// @Summary 创建ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonUserOrder true "创建ronUserOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ronUserOrder/createRonUserOrder [post]
export const createRonUserOrder = (data) => {
  return service({
    url: '/ronUserOrder/createRonUserOrder',
    method: 'post',
    data
  })
}

// @Tags RonUserOrder
// @Summary 删除ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonUserOrder true "删除ronUserOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ronUserOrder/deleteRonUserOrder [delete]
export const deleteRonUserOrder = (params) => {
  return service({
    url: '/ronUserOrder/deleteRonUserOrder',
    method: 'delete',
    params
  })
}

// @Tags RonUserOrder
// @Summary 批量删除ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ronUserOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ronUserOrder/deleteRonUserOrder [delete]
export const deleteRonUserOrderByIds = (params) => {
  return service({
    url: '/ronUserOrder/deleteRonUserOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags RonUserOrder
// @Summary 更新ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RonUserOrder true "更新ronUserOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ronUserOrder/updateRonUserOrder [put]
export const updateRonUserOrder = (data) => {
  return service({
    url: '/ronUserOrder/updateRonUserOrder',
    method: 'put',
    data
  })
}

// @Tags RonUserOrder
// @Summary 用id查询ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.RonUserOrder true "用id查询ronUserOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ronUserOrder/findRonUserOrder [get]
export const findRonUserOrder = (params) => {
  return service({
    url: '/ronUserOrder/findRonUserOrder',
    method: 'get',
    params
  })
}

// @Tags RonUserOrder
// @Summary 分页获取ronUserOrder表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ronUserOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ronUserOrder/getRonUserOrderList [get]
export const getRonUserOrderList = (params) => {
  return service({
    url: '/ronUserOrder/getRonUserOrderList',
    method: 'get',
    params
  })
}

// @Tags RonUserOrder
// @Summary 不需要鉴权的ronUserOrder表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.RonUserOrderSearch true "分页获取ronUserOrder表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ronUserOrder/getRonUserOrderPublic [get]
export const getRonUserOrderPublic = () => {
  return service({
    url: '/ronUserOrder/getRonUserOrderPublic',
    method: 'get',
  })
}
