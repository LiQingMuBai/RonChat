package system

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/zegoim/zego_server_assistant/token/go/src/token04"
	"go.uber.org/zap"
)

func (b *BaseApi) GetAuthToken(c *gin.Context) {

	var r systemReq.ZEGO
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var appId uint32 = 238914567                       // Zego派发的数字ID, 各个开发者的唯一标识
	userId := r.UserId                                 // 用户 ID
	serverSecret := "fa94dd0f974cf2e293728a526b028271" // 在获取 token 时进行 AES 加密的密钥
	var effectiveTimeInSeconds int64 = 3600            // token 的有效时长，单位：秒
	var payload string = ""                            // token业务认证扩展，基础鉴权token此处填空
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

	var r systemReq.ZEGO
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var appId uint32 = 238914567                       // Zego派发的数字ID, 各个开发者的唯一标识
	roomId := r.Roomid                                 // 房间 ID
	userId := r.UserId                                 // 用户 ID
	serverSecret := "0f9632b53ff717617bdb33989b633bc4" // 在获取 token 时进行 AES 加密的密钥
	var effectiveTimeInSeconds int64 = 3600            // token 的有效时长，单位：秒
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
