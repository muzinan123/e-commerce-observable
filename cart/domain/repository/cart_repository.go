package repository

import (
	"errors"

	"git.imooc.com/coding-447/cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (u *CartRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (u *CartRepository) FindCartByID(cartID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, u.mysqlDb.First(cart, cartID).Error
}

func (u *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := u.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("xxxx")
	}
	return cart.ID, nil
}

func (u *CartRepository) DeleteCartByID(cartID int64) error {
	return u.mysqlDb.Where("id = ?", cartID).Delete(&model.Cart{}).Error
}

func (u *CartRepository) UpdateCart(cart *model.Cart) error {
	return u.mysqlDb.Model(cart).Update(cart).Error
}

func (u *CartRepository) FindAll(userID int64) (cartAll []model.Cart, err error) {
	return cartAll, u.mysqlDb.Where("user_id = ?", userID).Find(&cartAll).Error
}

func (u *CartRepository) CleanCart(userID int64) error {
	return u.mysqlDb.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}

func (u *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (u *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := u.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
