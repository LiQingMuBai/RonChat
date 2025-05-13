// 自动生成模板RonUsers
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ronUsers表 结构体  RonUsers
type RonUsers struct {
	global.GVA_MODEL
	Username  string `json:"username" form:"username" gorm:"comment:用户登录名;column:username;size:191;"`      //用户登录名
	Password  string `json:"password" form:"password" gorm:"comment:用户登录密码;column:password;size:191;"`     //用户登录密码
	Sex       int    `json:"sex" form:"sex" gorm:"comment:男女 1男 0女;column:sex;size:19;"`                   //男女 1男 0女
	Enable    int    `json:"enable" form:"enable" gorm:"comment:用户是否被冻结 1正常 2冻结;column:enable;size:19;"`   //用户是否被冻结 1正常 2冻结
	HeaderImg string `json:"headerImg" form:"headerImg" gorm:"comment:用户头像;column:header_img;size:191;"`   //用户头像
	Phone     string `json:"phone" form:"phone" gorm:"comment:用户手机号;column:phone;size:191;"`               //用户手机号
	Email     string `json:"email" form:"email" gorm:"comment:用户邮箱;column:email;size:191;"`                //用户邮箱
	Telegram  string `json:"telegram" form:"telegram" gorm:"comment:用户telegram;column:telegram;size:191;"` //用户telegram
	RoomId    string `json:"roomId" form:"roomId" gorm:"comment:roomID;column:room_id;size:191;"`          //roomID
	Balance   string `json:"balance" form:"balance" gorm:"comment:balance;column:balance;size:191;"`       //balance
	Anchor    int    `json:"anchor" form:"anchor" gorm:"comment: 1主播 0非主播;column:anchor;size:19;"`         // 1主播 0非主播
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (RonUsers) TableName() string {
	return "ron_users"
}
