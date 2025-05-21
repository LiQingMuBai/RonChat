package system

import (
	"github.com/gin-gonic/gin"
)

type BaseZEGORouter struct{}

func (s *BaseRouter) InitBaseZEGORouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("zego")
	{
		baseRouter.POST("getAuthToken", baseApi.GetAuthToken)
		baseRouter.POST("login", baseApi.LoginRon)
		baseRouter.POST("getPermissionToken", baseApi.GetPermissionToken)

		baseRouter.POST("/create-checkout-session", baseApi.CreateCheckoutSession)
		baseRouter.POST("/payment/callback", baseApi.HandleStripePaymentWebhook)
	}
	return baseRouter
}
