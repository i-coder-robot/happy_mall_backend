package repository

import (
	"fmt"

	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/utils"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	List(req *query.ListQuery) (Orders []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(Order model.Order) (*model.Order, error)
	Exist(Order model.Order) *model.Order
	ExistByOrderID(id string) *model.Order
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(u model.Order) (bool, error)
}

func (repo *OrderRepository) List(req *query.ListQuery) (order []*model.Order, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.PageSize, req.Page) // 分页

	if err := db.Limit(limit).Offset(offset).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *OrderRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var orders []model.Order
	db := repo.DB

	if err := db.Find(&orders).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *OrderRepository) Get(order model.Order) (*model.Order, error) {
	if err := repo.DB.Where(&order).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *OrderRepository) Exist(order model.Order) *model.Order {
	if order.OrderId != "" {
		repo.DB.Model(&order).Where("order_id= ?", order.OrderId)
		return &order
	}
	return nil
}

func (repo *OrderRepository) ExistByOrderID(id string) *model.Order {
	var order model.Order
	repo.DB.Where("order_id = ?", id).First(&order)
	return &order
}

func (repo *OrderRepository) Add(order model.Order) (*model.Order, error) {
	err := repo.DB.Create(order).Error
	if err != nil {
		return &order, fmt.Errorf("订单添加失败")
	}
	return &order, nil
}

func (repo *OrderRepository) Edit(order model.Order) (bool, error) {
	if order.OrderId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	o := &model.Order{}
	err := repo.DB.Model(o).Where("order_id=?", order.OrderId).Updates(map[string]interface{}{
		"nick_name":    order.NickName,
		"mobile":       order.Mobile,
		"pay_status":   order.PayStatus,
		"order_status": order.OrderStatus,
		"extra_info":   order.ExtraInfo,
		"user_address": order.UserAddress,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *OrderRepository) Delete(o model.Order) (bool, error) {
	err := repo.DB.Model(&o).Where("order_id=?", o.OrderId).Update("is_deleted", o.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
