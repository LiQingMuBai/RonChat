// 自动生成模板RonUserOrder
package system

import (
	"time"
)

// ronUserOrder表 结构体  RonUserOrder
type RonUserOrder struct {
	Id        int       `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`                                //id字段
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                             //createdAt字段
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                             //updatedAt字段
	DeletedAt time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`                             //deletedAt字段
	UserId    uint      `json:"userId" form:"userId" gorm:"comment:用户登录密码;column:user_id;size:19;"`         //用户登录密码
	Status    int       `json:"status" form:"status" gorm:"comment:0 过程中 1 成功 2失败;column:status;size:19;"` //0 过程中 1 成功 2失败
	Amount    string    `json:"amount" form:"amount" gorm:"comment:充值金额;column:amount;size:191;"`             //充值金额
	Channel   string    `json:"channel" form:"channel" gorm:"comment:充值通道;column:channel;size:191;"`          //充值通道
	OrderId   string    `json:"order_id" form:"channel" gorm:"comment:订单id;column:channel;size:191;"`           //订单id
}

// TableName ronUserOrder表 RonUserOrder自定义表名 ron_user_order
func (RonUserOrder) TableName() string {
	return "ron_user_order"
}
