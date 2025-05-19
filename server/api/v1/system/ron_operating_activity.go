package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type RonOperatingActivityApi struct {}



// CreateRonOperatingActivity 创建ronOperatingActivity表
// @Tags RonOperatingActivity
// @Summary 创建ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonOperatingActivity true "创建ronOperatingActivity表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ronOperatingActivity/createRonOperatingActivity [post]
func (ronOperatingActivityApi *RonOperatingActivityApi) CreateRonOperatingActivity(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var ronOperatingActivity system.RonOperatingActivity
	err := c.ShouldBindJSON(&ronOperatingActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ronOperatingActivityService.CreateRonOperatingActivity(ctx,&ronOperatingActivity)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteRonOperatingActivity 删除ronOperatingActivity表
// @Tags RonOperatingActivity
// @Summary 删除ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonOperatingActivity true "删除ronOperatingActivity表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ronOperatingActivity/deleteRonOperatingActivity [delete]
func (ronOperatingActivityApi *RonOperatingActivityApi) DeleteRonOperatingActivity(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := ronOperatingActivityService.DeleteRonOperatingActivity(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRonOperatingActivityByIds 批量删除ronOperatingActivity表
// @Tags RonOperatingActivity
// @Summary 批量删除ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ronOperatingActivity/deleteRonOperatingActivityByIds [delete]
func (ronOperatingActivityApi *RonOperatingActivityApi) DeleteRonOperatingActivityByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := ronOperatingActivityService.DeleteRonOperatingActivityByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRonOperatingActivity 更新ronOperatingActivity表
// @Tags RonOperatingActivity
// @Summary 更新ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonOperatingActivity true "更新ronOperatingActivity表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ronOperatingActivity/updateRonOperatingActivity [put]
func (ronOperatingActivityApi *RonOperatingActivityApi) UpdateRonOperatingActivity(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var ronOperatingActivity system.RonOperatingActivity
	err := c.ShouldBindJSON(&ronOperatingActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ronOperatingActivityService.UpdateRonOperatingActivity(ctx,ronOperatingActivity)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRonOperatingActivity 用id查询ronOperatingActivity表
// @Tags RonOperatingActivity
// @Summary 用id查询ronOperatingActivity表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询ronOperatingActivity表"
// @Success 200 {object} response.Response{data=system.RonOperatingActivity,msg=string} "查询成功"
// @Router /ronOperatingActivity/findRonOperatingActivity [get]
func (ronOperatingActivityApi *RonOperatingActivityApi) FindRonOperatingActivity(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reronOperatingActivity, err := ronOperatingActivityService.GetRonOperatingActivity(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reronOperatingActivity, c)
}
// GetRonOperatingActivityList 分页获取ronOperatingActivity表列表
// @Tags RonOperatingActivity
// @Summary 分页获取ronOperatingActivity表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.RonOperatingActivitySearch true "分页获取ronOperatingActivity表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ronOperatingActivity/getRonOperatingActivityList [get]
func (ronOperatingActivityApi *RonOperatingActivityApi) GetRonOperatingActivityList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.RonOperatingActivitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ronOperatingActivityService.GetRonOperatingActivityInfoList(ctx,pageInfo)
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

// GetRonOperatingActivityPublic 不需要鉴权的ronOperatingActivity表接口
// @Tags RonOperatingActivity
// @Summary 不需要鉴权的ronOperatingActivity表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ronOperatingActivity/getRonOperatingActivityPublic [get]
func (ronOperatingActivityApi *RonOperatingActivityApi) GetRonOperatingActivityPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    ronOperatingActivityService.GetRonOperatingActivityPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的ronOperatingActivity表接口信息",
    }, "获取成功", c)
}
