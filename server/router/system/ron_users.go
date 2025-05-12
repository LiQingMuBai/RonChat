package system

import (
	"github.com/gin-gonic/gin"
)

type RonUsersRouter struct{}

// InitRonUsersRouter 初始化 ronUsers表 路由信息
func (s *RonUsersRouter) InitRonUsersRouter(PublicRouter *gin.RouterGroup) {

	ronUsersRouterWithoutAuth := PublicRouter.Group("ronUsers")
	{
		ronUsersRouterWithoutAuth.POST("createRonUsers", ronUsersApi.CreateRonUsers)             // 新建ronUsers表
		ronUsersRouterWithoutAuth.DELETE("deleteRonUsers", ronUsersApi.DeleteRonUsers)           // 删除ronUsers表
		ronUsersRouterWithoutAuth.DELETE("deleteRonUsersByIds", ronUsersApi.DeleteRonUsersByIds) // 批量删除ronUsers表
		ronUsersRouterWithoutAuth.PUT("updateRonUsers", ronUsersApi.UpdateRonUsers)              // 更新ronUsers表
		ronUsersRouterWithoutAuth.GET("findRonUsers", ronUsersApi.FindRonUsers)                  // 根据ID获取ronUsers表
		ronUsersRouterWithoutAuth.GET("findRonUserByTG", ronUsersApi.FindRonUserByTG)            // 根据ID获取ronUsers表
		ronUsersRouterWithoutAuth.GET("getRonUsersList", ronUsersApi.GetRonUsersList)            // 获取ronUsers表列表
		ronUsersRouterWithoutAuth.GET("getRonUsersPublic", ronUsersApi.GetRonUsersPublic)        // ronUsers表开放接口

		ronUsersRouterWithoutAuth.PUT("enterRoom", ronUsersApi.UpdateRonUsers) // 进入房间
		ronUsersRouterWithoutAuth.PUT("exitRoom", ronUsersApi.UpdateRonUsers)  // 退出房间

	}
}
