package system

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
	"github.com/stripe/stripe-go/v82/webhook"
	"github.com/zegoim/zego_server_assistant/token/go/src/token04"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (b *BaseApi) LoginRon(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user, err := userService.FindUserBy3rdID(l.Username)
	if err != nil {
		user := &system.SysUser{
			Username:    l.Username,
			NickName:    l.Username,
			Password:    "123456",
			AuthorityId: 999, Enable: 1}
		newUser, _ := userService.Register(*user)

		var ronUser system.RonUsers
		ronUser.Uid = newUser.ID
		ronUser.Username = l.Username
		ronUser.Telegram = l.Username
		ronUser.Enable = 0

		err = ronUsersService.CreateRonUsers(c, &ronUser)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}

		b.TokenNext(c, newUser)
		return
	} else {
		if user.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")

			response.FailWithMessage("用户被禁止登录", c)
			return
		}
	}
	b.TokenNext(c, *user)
	return

}

func stringToUint32(input string) uint32 {
	// First convert to uint64 to check for overflow
	val, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(val)
}
func (b *BaseApi) GetAuthToken(c *gin.Context) {

	uid := utils.GetUserID(c)
	var r systemReq.ZEGO
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var appId uint32 = stringToUint32(global.GVA_CONFIG.System.AppID) // Zego派发的数字ID, 各个开发者的唯一标识
	log.Println(appId)
	userId := fmt.Sprintf("%d", uid)

	log.Println("userId:", userId)
	// 用户 ID
	serverSecret := global.GVA_CONFIG.System.ServerSecret // 在获取 token 时进行 AES 加密的密钥
	log.Println(serverSecret)
	var effectiveTimeInSeconds int64 = 3600 // token 的有效时长，单位：秒
	var payload string = ""                 // token业务认证扩展，基础鉴权token此处填空
	//生成token
	token, err := token04.GenerateToken04(appId, userId, serverSecret, effectiveTimeInSeconds, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)

	if err != nil {
		global.GVA_LOG.Error("验证Token获取失败!", zap.Error(err))
		response.FailWithMessage("验证Token获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.ZEGOResponse{
		Token: token,
	}, "验证Token获取成功", c)
}
func (b *BaseApi) GetPermissionToken(c *gin.Context) {
	uid := utils.GetUserID(c)
	var r systemReq.ZEGO
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var appId uint32 = stringToUint32(global.GVA_CONFIG.System.AppID) // Zego派发的数字ID, 各个开发者的唯一标识
	roomId := r.Roomid                                                // 房间 ID
	userId := fmt.Sprintf("%d", uid)                                  // 用户 ID
	serverSecret := global.GVA_CONFIG.System.ServerSecret             // 在获取 token 时进行 AES 加密的密钥
	var effectiveTimeInSeconds int64 = 3600                           // token 的有效时长，单位：秒
	//请参考 github.com/zegoim/zego_server_assistant/token/go/src/token04/token04.go 定义
	////权限位定义
	//const (
	//    PrivilegeKeyLogin   = 1 // 登录权限位认证
	//    PrivilegeKeyPublish = 2 // 推流权限位认证
	//)
	////权限开关定义
	//const (
	//    PrivilegeEnable     = 1 // 有权限
	//    PrivilegeDisable    = 0 // 无权限
	//)
	//业务权限认证配置，可以配置多个权限位
	privilege := make(map[int]int)
	privilege[token04.PrivilegeKeyLogin] = token04.PrivilegeEnable   // 有房间登录权限
	privilege[token04.PrivilegeKeyPublish] = token04.PrivilegeEnable // 无推流权限
	//token业务扩展配置
	payloadData := &RtcRoomPayLoad{
		RoomId:       roomId,
		Privilege:    privilege,
		StreamIdList: nil,
	}
	payload, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println(err)
		return
	}
	//生成token
	token, err := token04.GenerateToken04(appId, userId, serverSecret, effectiveTimeInSeconds, string(payload))
	if err != nil {
		global.GVA_LOG.Error("权限Token获取失败!", zap.Error(err))
		response.FailWithMessage("权限Token获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.ZEGOResponse{
		Token: token,
	}, "权限Token获取成功", c)
}

/*
权限认证token生成示例代码
*/
//token业务扩展：权限认证属性
type RtcRoomPayLoad struct {
	RoomId       string      `json:"room_id"`        //房间 id（必填）；用于对接口的房间 id 进行强验证
	Privilege    map[int]int `json:"privilege"`      //权限位开关列表；用于对接口的操作权限进行强验证
	StreamIdList []string    `json:"stream_id_list"` //流列表；用于对接口的流 id 进行强验证；允许为空，如果为空，则不对流 id 验证
}

func (b *BaseApi) CreateCheckoutSession(c *gin.Context) {
	uid := utils.GetUserID(c)
	userId := fmt.Sprintf("%d", uid)

	orderId := utils.GenerateOrderID(userId)

	var ronUserOrder system.RonUserOrder
	err := c.ShouldBindJSON(&ronUserOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(2000), // 金额，单位为分（例如 20.00 USD）
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),

		// 添加自定义参数
		Metadata: map[string]string{
			"user_id":  userId,
			"order_id": orderId,
			"env":      "sandbox",
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		global.GVA_LOG.Error("创建支付意图失败", zap.Error(err))
		response.FailWithMessage("创建支付意图失败", c)
		return
	}

	//var ronUserOrder system.RonUserOrder

	ronUserOrder.UserId = uid
	ronUserOrder.Status = 0
	ronUserOrder.OrderId = orderId

	err = ronUserOrderService.CreateRonUserOrder(c, &ronUserOrder)
	if err != nil {
		global.GVA_LOG.Error("创建支付意图失败", zap.Error(err))
		response.FailWithMessage("创建支付意图失败", c)
		return
	}

	response.OkWithDetailed(systemRes.StripeResponse{
		ClientSecret: pi.ClientSecret,
	}, "创建支付意图成功", c)

}

// HandleWebhook 处理 Stripe Webhook 事件
func (b *BaseApi) HandleStripePaymentWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// 替换为你的 Webhook Secret
	endpointSecret := "whsec_Q9lhMsepl1vOrhxzvkLmAiggn3D08MR9"

	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		log.Printf("Webhook signature verification failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Webhook signature verification failed"})
		return
	}

	// 处理不同的事件类型
	switch event.Type {
	case "payment_intent.succeeded":
		// 提取 PaymentIntent 对象
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			log.Printf("Failed to parse payment intent: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse payment intent"})
			return
		}

		// 获取 metadata 中的自定义参数
		userID := paymentIntent.Metadata["user_id"]
		orderID := paymentIntent.Metadata["order_id"]
		env := paymentIntent.Metadata["env"]

		log.Printf("Payment succeeded: user_id=%s, order_id=%s, env=%s", userID, orderID, env)
		//订单完成
		var ronUserOrder system.RonUserOrder
		ronUserOrder.Status = 1
		ronUserOrder.OrderId = orderID
		err = ronUserOrderService.UpdateRonUserOrderByOrderID(c, ronUserOrder)

		//客户余额充值
		var ronUsers system.RonUsers
		_userId, _ := strconv.ParseUint(userID, 10, 64)
		record, _ := ronUsersService.GetRonUsersByUID(c, _userId)
		balance, _ := utils.AddDecimalStringsWithPrecision(record.Balance, ronUserOrder.Amount, 2)
		record.Balance = balance
		err = ronUsersService.UpdateRonUsers(c, ronUsers)
		if err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败:"+err.Error(), c)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "Payment succeeded",
			"user_id":  userID,
			"order_id": orderID,
			"env":      env,
		})

	case "payment_intent.payment_failed":
		// 提取 PaymentIntent 对象
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			log.Printf("Failed to parse payment intent: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse payment intent"})
			return
		}

		// 获取 metadata 中的自定义参数
		userID := paymentIntent.Metadata["user_id"]
		orderID := paymentIntent.Metadata["order_id"]
		env := paymentIntent.Metadata["env"]

		log.Printf("Payment failed: user_id=%s, order_id=%s, env=%s", userID, orderID, env)

		var ronUserOrder system.RonUserOrder
		ronUserOrder.Status = 2
		ronUserOrder.OrderId = orderID
		err = ronUserOrderService.UpdateRonUserOrderByOrderID(c, ronUserOrder)

		c.JSON(http.StatusOK, gin.H{
			"status":   "Payment failed",
			"user_id":  userID,
			"order_id": orderID,
			"env":      env,
		})

	default:
		log.Printf("Unhandled event type: %s", event.Type)
		c.JSON(http.StatusOK, gin.H{"status": "Unhandled event type"})
	}
}
