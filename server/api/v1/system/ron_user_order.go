package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type RonUserOrderApi struct {}



// CreateRonUserOrder 创建ronUserOrder表
// @Tags RonUserOrder
// @Summary 创建ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonUserOrder true "创建ronUserOrder表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ronUserOrder/createRonUserOrder [post]
func (ronUserOrderApi *RonUserOrderApi) CreateRonUserOrder(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var ronUserOrder system.RonUserOrder
	err := c.ShouldBindJSON(&ronUserOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ronUserOrderService.CreateRonUserOrder(ctx,&ronUserOrder)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteRonUserOrder 删除ronUserOrder表
// @Tags RonUserOrder
// @Summary 删除ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonUserOrder true "删除ronUserOrder表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ronUserOrder/deleteRonUserOrder [delete]
func (ronUserOrderApi *RonUserOrderApi) DeleteRonUserOrder(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := ronUserOrderService.DeleteRonUserOrder(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRonUserOrderByIds 批量删除ronUserOrder表
// @Tags RonUserOrder
// @Summary 批量删除ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ronUserOrder/deleteRonUserOrderByIds [delete]
func (ronUserOrderApi *RonUserOrderApi) DeleteRonUserOrderByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := ronUserOrderService.DeleteRonUserOrderByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRonUserOrder 更新ronUserOrder表
// @Tags RonUserOrder
// @Summary 更新ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonUserOrder true "更新ronUserOrder表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ronUserOrder/updateRonUserOrder [put]
func (ronUserOrderApi *RonUserOrderApi) UpdateRonUserOrder(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var ronUserOrder system.RonUserOrder
	err := c.ShouldBindJSON(&ronUserOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ronUserOrderService.UpdateRonUserOrder(ctx,ronUserOrder)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRonUserOrder 用id查询ronUserOrder表
// @Tags RonUserOrder
// @Summary 用id查询ronUserOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询ronUserOrder表"
// @Success 200 {object} response.Response{data=system.RonUserOrder,msg=string} "查询成功"
// @Router /ronUserOrder/findRonUserOrder [get]
func (ronUserOrderApi *RonUserOrderApi) FindRonUserOrder(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reronUserOrder, err := ronUserOrderService.GetRonUserOrder(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reronUserOrder, c)
}
// GetRonUserOrderList 分页获取ronUserOrder表列表
// @Tags RonUserOrder
// @Summary 分页获取ronUserOrder表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.RonUserOrderSearch true "分页获取ronUserOrder表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ronUserOrder/getRonUserOrderList [get]
func (ronUserOrderApi *RonUserOrderApi) GetRonUserOrderList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.RonUserOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ronUserOrderService.GetRonUserOrderInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetRonUserOrderPublic 不需要鉴权的ronUserOrder表接口
// @Tags RonUserOrder
// @Summary 不需要鉴权的ronUserOrder表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ronUserOrder/getRonUserOrderPublic [get]
func (ronUserOrderApi *RonUserOrderApi) GetRonUserOrderPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    ronUserOrderService.GetRonUserOrderPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的ronUserOrder表接口信息",
    }, "获取成功", c)
}
