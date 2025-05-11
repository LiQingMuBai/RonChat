package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type RonUsersService struct{}

// CreateRonUsers 创建ronUsers表记录
// Author [yourname](https://github.com/yourname)
func (ronUsersService *RonUsersService) CreateRonUsers(ctx context.Context, ronUsers *system.RonUsers) (err error) {
	err = global.GVA_DB.Create(ronUsers).Error
	return err
}

// DeleteRonUsers 删除ronUsers表记录
// Author [yourname](https://github.com/yourname)
func (ronUsersService *RonUsersService) DeleteRonUsers(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&system.RonUsers{}, "id = ?", id).Error
	return err
}

// DeleteRonUsersByIds 批量删除ronUsers表记录
// Author [yourname](https://github.com/yourname)
func (ronUsersService *RonUsersService) DeleteRonUsersByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.RonUsers{}, "id in ?", ids).Error
	return err
}

// UpdateRonUsers 更新ronUsers表记录
// Author [yourname](https://github.com/yourname)
func (ronUsersService *RonUsersService) UpdateRonUsers(ctx context.Context, ronUsers system.RonUsers) (err error) {
	err = global.GVA_DB.Model(&system.RonUsers{}).Where("id = ?", ronUsers.ID).Updates(&ronUsers).Error
	return err
}

// GetRonUsers 根据id获取ronUsers表记录
// Author [yourname](https://github.com/yourname)
func (ronUsersService *RonUsersService) GetRonUsers(ctx context.Context, id string) (ronUsers system.RonUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ronUsers).Error
	return
}

// GetRonUsersInfoList 分页获取ronUsers表记录
// Author [yourname](https://github.com/yourname)
func (ronUsersService *RonUsersService) GetRonUsersInfoList(ctx context.Context, info systemReq.RonUsersSearch) (list []system.RonUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.RonUsers{})
	var ronUserss []system.RonUsers
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ronUserss).Error
	return ronUserss, total, err
}
func (ronUsersService *RonUsersService) GetRonUsersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
