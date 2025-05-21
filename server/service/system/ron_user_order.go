package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type RonUserOrderService struct{}

// CreateRonUserOrder 创建ronUserOrder表记录
// Author [yourname](https://github.com/yourname)
func (ronUserOrderService *RonUserOrderService) CreateRonUserOrder(ctx context.Context, ronUserOrder *system.RonUserOrder) (err error) {
	err = global.GVA_DB.Create(ronUserOrder).Error
	return err
}

// DeleteRonUserOrder 删除ronUserOrder表记录
// Author [yourname](https://github.com/yourname)
func (ronUserOrderService *RonUserOrderService) DeleteRonUserOrder(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&system.RonUserOrder{}, "id = ?", id).Error
	return err
}

// DeleteRonUserOrderByIds 批量删除ronUserOrder表记录
// Author [yourname](https://github.com/yourname)
func (ronUserOrderService *RonUserOrderService) DeleteRonUserOrderByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.RonUserOrder{}, "id in ?", ids).Error
	return err
}

// UpdateRonUserOrder 更新ronUserOrder表记录
// Author [yourname](https://github.com/yourname)
func (ronUserOrderService *RonUserOrderService) UpdateRonUserOrder(ctx context.Context, ronUserOrder system.RonUserOrder) (err error) {
	err = global.GVA_DB.Model(&system.RonUserOrder{}).Where("id = ?", ronUserOrder.Id).Updates(&ronUserOrder).Error
	return err
}

// UpdateRonUserOrder 更新ronUserOrder表记录
// Author [yourname](https://github.com/yourname)
func (ronUserOrderService *RonUserOrderService) UpdateRonUserOrderByOrderID(ctx context.Context, ronUserOrder system.RonUserOrder) (err error) {
	err = global.GVA_DB.Model(&system.RonUserOrder{}).Where("order_id = ?", ronUserOrder.OrderId).Updates(&ronUserOrder).Error
	return err
}

// GetRonUserOrder 根据id获取ronUserOrder表记录
// Author [yourname](https://github.com/yourname)
func (ronUserOrderService *RonUserOrderService) GetRonUserOrder(ctx context.Context, id string) (ronUserOrder system.RonUserOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ronUserOrder).Error
	return
}

// GetRonUserOrderInfoList 分页获取ronUserOrder表记录
// Author [yourname](https://github.com/yourname)
func (ronUserOrderService *RonUserOrderService) GetRonUserOrderInfoList(ctx context.Context, info systemReq.RonUserOrderSearch) (list []system.RonUserOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.RonUserOrder{})
	var ronUserOrders []system.RonUserOrder
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ronUserOrders).Error
	return ronUserOrders, total, err
}
func (ronUserOrderService *RonUserOrderService) GetRonUserOrderPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
