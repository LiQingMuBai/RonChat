package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"strconv"
)

type RonUsersApi struct{}

// CreateRonUsers 创建ronUsers表
// @Tags RonUsers
// @Summary 创建ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonUsers true "创建ronUsers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ronUsers/createRonUsers [post]
func (ronUsersApi *RonUsersApi) CreateRonUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ronUsers system.RonUsers
	err := c.ShouldBindJSON(&ronUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ronUsersService.CreateRonUsers(ctx, &ronUsers)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithData(ronUsers, c)
}

// DeleteRonUsers 删除ronUsers表
// @Tags RonUsers
// @Summary 删除ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonUsers true "删除ronUsers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ronUsers/deleteRonUsers [delete]
func (ronUsersApi *RonUsersApi) DeleteRonUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := ronUsersService.DeleteRonUsers(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRonUsersByIds 批量删除ronUsers表
// @Tags RonUsers
// @Summary 批量删除ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ronUsers/deleteRonUsersByIds [delete]
func (ronUsersApi *RonUsersApi) DeleteRonUsersByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := ronUsersService.DeleteRonUsersByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRonUsers 更新ronUsers表
// @Tags RonUsers
// @Summary 更新ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.RonUsers true "更新ronUsers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ronUsers/updateRonUsers [put]
func (ronUsersApi *RonUsersApi) UpdateRonUsers(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ronUsers system.RonUsers
	err := c.ShouldBindJSON(&ronUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ronUsersService.UpdateRonUsers(ctx, ronUsers)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
func (ronUsersApi *RonUsersApi) CreateRoom(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ronUsers system.RonUsers
	err := c.ShouldBindJSON(&ronUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ronUsers.Enable = 1

	err = ronUsersService.UpdateRonUsers(ctx, ronUsers)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
func (ronUsersApi *RonUsersApi) ExitRoom(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ronUsers system.RonUsers
	err := c.ShouldBindJSON(&ronUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ronUsers.Enable = 0
	ronUsers.StreamId = ""

	err = ronUsersService.UpdateRonUsers(ctx, ronUsers)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRonUsers 用id查询ronUsers表
// @Tags RonUsers
// @Summary 用id查询ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询ronUsers表"
// @Success 200 {object} response.Response{data=system.RonUsers,msg=string} "查询成功"
// @Router /ronUsers/findRonUsers [get]
func (ronUsersApi *RonUsersApi) FindRonUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reronUsers, err := ronUsersService.GetRonUsers(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reronUsers, c)
}

// FindRonUsers 用id查询ronUsers表
// @Tags RonUsers
// @Summary 用id查询ronUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询ronUsers表"
// @Success 200 {object} response.Response{data=system.RonUsers,msg=string} "查询成功"
// @Router /ronUsers/findRonUsers [get]
func (ronUsersApi *RonUsersApi) FindRonUserByTG(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("telegram")
	reronUsers, err := ronUsersService.GetRonUserByTG(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reronUsers, c)
}

func (ronUsersApi *RonUsersApi) GetLuckyGuy(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	sex := c.Query("sex")

	isMale, err := strconv.Atoi(sex)

	if err != nil {
		isMale = 0
	}
	uid := utils.GetUserID(c)
	log.Println("uid : ", uid)
	log.Println("isMale : ", isMale)

	ronUsers, max, err := ronUsersService.GetRonUsersPublic(ctx, isMale, uid)

	log.Println("max :  ", max)
	if err != nil {
		global.GVA_LOG.Error("当前用户缺少!", zap.Error(err))
		response.FailWithMessage("当前用户缺少:"+err.Error(), c)
		return
	}

	if max <= 1 {
		global.GVA_LOG.Error("当前用户缺少!", zap.Error(err))
		response.OkWithCode(c)
	}
	randomNum := rand.Intn(int(max-0)) + 0
	response.OkWithData(ronUsers[randomNum], c)
}

// GetRonUsersList 分页获取ronUsers表列表
// @Tags RonUsers
// @Summary 分页获取ronUsers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.RonUsersSearch true "分页获取ronUsers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ronUsers/getRonUsersList [get]
func (ronUsersApi *RonUsersApi) GetRonUsersList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.RonUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ronUsersService.GetRonUsersInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

//// GetRonUsersPublic 不需要鉴权的ronUsers表接口
//// @Tags RonUsers
//// @Summary 不需要鉴权的ronUsers表接口
//// @Accept application/json
//// @Produce application/json
//// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
//// @Router /ronUsers/getRonUsersPublic [get]
//func (ronUsersApi *RonUsersApi) GetRonUsersPublic(c *gin.Context) {
//	// 创建业务用Context
//	ctx := c.Request.Context()
//
//	// 此接口不需要鉴权
//	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
//	ronUsersService.GetRonUsersPublic(ctx, 1,1)
//	response.OkWithDetailed(gin.H{
//		"info": "不需要鉴权的ronUsers表接口信息",
//	}, "获取成功", c)
//}
