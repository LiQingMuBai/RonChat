
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type RonOperatingActivityService struct {}
// CreateRonOperatingActivity 创建ronOperatingActivity表记录
// Author [yourname](https://github.com/yourname)
func (ronOperatingActivityService *RonOperatingActivityService) CreateRonOperatingActivity(ctx context.Context, ronOperatingActivity *system.RonOperatingActivity) (err error) {
	err = global.GVA_DB.Create(ronOperatingActivity).Error
	return err
}

// DeleteRonOperatingActivity 删除ronOperatingActivity表记录
// Author [yourname](https://github.com/yourname)
func (ronOperatingActivityService *RonOperatingActivityService)DeleteRonOperatingActivity(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&system.RonOperatingActivity{},"id = ?",id).Error
	return err
}

// DeleteRonOperatingActivityByIds 批量删除ronOperatingActivity表记录
// Author [yourname](https://github.com/yourname)
func (ronOperatingActivityService *RonOperatingActivityService)DeleteRonOperatingActivityByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.RonOperatingActivity{},"id in ?",ids).Error
	return err
}

// UpdateRonOperatingActivity 更新ronOperatingActivity表记录
// Author [yourname](https://github.com/yourname)
func (ronOperatingActivityService *RonOperatingActivityService)UpdateRonOperatingActivity(ctx context.Context, ronOperatingActivity system.RonOperatingActivity) (err error) {
	err = global.GVA_DB.Model(&system.RonOperatingActivity{}).Where("id = ?",ronOperatingActivity.Id).Updates(&ronOperatingActivity).Error
	return err
}

// GetRonOperatingActivity 根据id获取ronOperatingActivity表记录
// Author [yourname](https://github.com/yourname)
func (ronOperatingActivityService *RonOperatingActivityService)GetRonOperatingActivity(ctx context.Context, id string) (ronOperatingActivity system.RonOperatingActivity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ronOperatingActivity).Error
	return
}
// GetRonOperatingActivityInfoList 分页获取ronOperatingActivity表记录
// Author [yourname](https://github.com/yourname)
func (ronOperatingActivityService *RonOperatingActivityService)GetRonOperatingActivityInfoList(ctx context.Context, info systemReq.RonOperatingActivitySearch) (list []system.RonOperatingActivity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&system.RonOperatingActivity{})
    var ronOperatingActivitys []system.RonOperatingActivity
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&ronOperatingActivitys).Error
	return  ronOperatingActivitys, total, err
}
func (ronOperatingActivityService *RonOperatingActivityService)GetRonOperatingActivityPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
