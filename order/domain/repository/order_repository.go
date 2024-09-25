package repository

import (
	"errors"

	"git.imooc.com/coding-447/order/domain/model"
	"github.com/jinzhu/gorm"
)

type IOrderRepository interface {
	InitTable() error
	FindOrderByID(int64) (*model.Order, error)
	CreateOrder(*model.Order) (int64, error)
	DeleteOrderByID(int64) error
	UpdateOrder(*model.Order) error
	FindAll() ([]model.Order, error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{mysqlDb: db}
}

type OrderRepository struct {
	mysqlDb *gorm.DB
}

func (u *OrderRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Order{}, &model.OrderDetail{}).Error
}

func (u *OrderRepository) FindOrderByID(orderID int64) (order *model.Order, err error) {
	order = &model.Order{}
	return order, u.mysqlDb.Preload("OrderDetail").First(order, orderID).Error
}

func (u *OrderRepository) CreateOrder(order *model.Order) (int64, error) {
	return order.ID, u.mysqlDb.Create(order).Error
}

func (u *OrderRepository) DeleteOrderByID(orderID int64) error {
	tx := u.mysqlDb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Unscoped().Where("id = ?", orderID).Delete(&model.Order{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("order_id = ?", orderID).Delete(&model.OrderDetail{}).Error; err != nil {
		tx.Rollback()
		return err

	}
	return tx.Commit().Error
}

func (u *OrderRepository) UpdateOrder(order *model.Order) error {
	return u.mysqlDb.Model(order).Update(order).Error
}

func (u *OrderRepository) FindAll() (orderAll []model.Order, err error) {
	return orderAll, u.mysqlDb.Preload("OrderDetail").Find(&orderAll).Error
}

func (u *OrderRepository) UpdateShipStatus(orderID int64, shipStatus int32) error {
	db := u.mysqlDb.Model(&model.Order{}).Where("id = ?", orderID).UpdateColumn("ship_status", shipStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func (u *OrderRepository) UpdatePayStatus(orderID int64, payStatus int32) error {
	db := u.mysqlDb.Model(&model.Order{}).Where("id = ?", orderID).UpdateColumn("pay_status", payStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}
