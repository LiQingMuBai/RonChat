
// 自动生成模板RonOperatingActivity
package system
import (
	"time"
)

// ronOperatingActivity表 结构体  RonOperatingActivity
type RonOperatingActivity struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  Name  *string `json:"name" form:"name" gorm:"comment:活动;column:name;size:191;"`  //活动
  VirtualValue  *string `json:"virtualValue" form:"virtualValue" gorm:"comment:虚拟值;column:virtual_value;size:191;"`  //虚拟值
  RealValue  *string `json:"realValue" form:"realValue" gorm:"comment:真实值;column:real_value;size:191;"`  //真实值
  Status  *bool `json:"status" form:"status" gorm:"comment:状态;column:status;"`  //状态
  Desc  *string `json:"desc" form:"desc" gorm:"comment:描述;column:desc;size:191;"`  //描述
}


// TableName ronOperatingActivity表 RonOperatingActivity自定义表名 ron_operating_activity
func (RonOperatingActivity) TableName() string {
    return "ron_operating_activity"
}





