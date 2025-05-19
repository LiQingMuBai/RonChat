package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RonOperatingActivityRouter struct{}

// InitRonOperatingActivityRouter 初始化 ronOperatingActivity表 路由信息
func (s *RonOperatingActivityRouter) InitRonOperatingActivityRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ronOperatingActivityRouter := Router.Group("ronOperatingActivity").Use(middleware.OperationRecord())
	//ronOperatingActivityRouterWithoutRecord := Router.Group("ronOperatingActivity")
	ronOperatingActivityRouterWithoutAuth := PublicRouter.Group("ronOperatingActivity")
	{
		ronOperatingActivityRouter.POST("createRonOperatingActivity", ronOperatingActivityApi.CreateRonOperatingActivity)             // 新建ronOperatingActivity表
		ronOperatingActivityRouter.DELETE("deleteRonOperatingActivity", ronOperatingActivityApi.DeleteRonOperatingActivity)           // 删除ronOperatingActivity表
		ronOperatingActivityRouter.DELETE("deleteRonOperatingActivityByIds", ronOperatingActivityApi.DeleteRonOperatingActivityByIds) // 批量删除ronOperatingActivity表
		ronOperatingActivityRouter.PUT("updateRonOperatingActivity", ronOperatingActivityApi.UpdateRonOperatingActivity)              // 更新ronOperatingActivity表
	}
	{
		ronOperatingActivityRouterWithoutAuth.GET("findRonOperatingActivity", ronOperatingActivityApi.FindRonOperatingActivity)       // 根据ID获取ronOperatingActivity表
		ronOperatingActivityRouterWithoutAuth.GET("getRonOperatingActivityList", ronOperatingActivityApi.GetRonOperatingActivityList) // 获取ronOperatingActivity表列表
	}
	{
		ronOperatingActivityRouterWithoutAuth.GET("getRonOperatingActivityPublic", ronOperatingActivityApi.GetRonOperatingActivityPublic) // ronOperatingActivity表开放接口
	}
}
