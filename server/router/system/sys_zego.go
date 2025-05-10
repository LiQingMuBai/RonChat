package system

import (
	"github.com/gin-gonic/gin"
)

type BaseZEGORouter struct{}

func (s *BaseRouter) InitBaseZEGORouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("zego")
	{
		baseRouter.POST("getAuthToken", baseApi.GetAuthToken)
		baseRouter.POST("getPermissionToken", baseApi.GetPermissionToken)
	}
	return baseRouter
}
