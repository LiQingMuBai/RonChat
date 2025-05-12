package system

import (
	"github.com/gin-gonic/gin"
)

type RonUserOrderRouter struct{}

// InitRonUserOrderRouter 初始化 ronUserOrder表 路由信息
func (s *RonUserOrderRouter) InitRonUserOrderRouter(PublicRouter *gin.RouterGroup) {

	ronUserOrderRouterWithoutAuth := PublicRouter.Group("ronUserOrder")
	{
		ronUserOrderRouterWithoutAuth.POST("createRonUserOrder", ronUserOrderApi.CreateRonUserOrder)             // 新建ronUserOrder表
		ronUserOrderRouterWithoutAuth.DELETE("deleteRonUserOrder", ronUserOrderApi.DeleteRonUserOrder)           // 删除ronUserOrder表
		ronUserOrderRouterWithoutAuth.DELETE("deleteRonUserOrderByIds", ronUserOrderApi.DeleteRonUserOrderByIds) // 批量删除ronUserOrder表
		ronUserOrderRouterWithoutAuth.PUT("updateRonUserOrder", ronUserOrderApi.UpdateRonUserOrder)              // 更新ronUserOrder表
	}
	{
		ronUserOrderRouterWithoutAuth.GET("findRonUserOrder", ronUserOrderApi.FindRonUserOrder)       // 根据ID获取ronUserOrder表
		ronUserOrderRouterWithoutAuth.GET("getRonUserOrderList", ronUserOrderApi.GetRonUserOrderList) // 获取ronUserOrder表列表
	}
	{
		ronUserOrderRouterWithoutAuth.GET("getRonUserOrderPublic", ronUserOrderApi.GetRonUserOrderPublic) // ronUserOrder表开放接口
	}
}
