package repository

import (
	"git.imooc.com/coding-447/payment/domain/model"
	"github.com/jinzhu/gorm"
)

type IPaymentRepository interface {
	InitTable() error
	FindPaymentByID(int64) (*model.Payment, error)
	CreatePayment(*model.Payment) (int64, error)
	DeletePaymentByID(int64) error
	UpdatePayment(*model.Payment) error
	FindAll() ([]model.Payment, error)
}

func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{mysqlDb: db}
}

type PaymentRepository struct {
	mysqlDb *gorm.DB
}

func (u *PaymentRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Payment{}).Error
}

func (u *PaymentRepository) FindPaymentByID(paymentID int64) (payment *model.Payment, err error) {
	payment = &model.Payment{}
	return payment, u.mysqlDb.First(payment, paymentID).Error
}

func (u *PaymentRepository) CreatePayment(payment *model.Payment) (int64, error) {
	return payment.ID, u.mysqlDb.Create(payment).Error
}

func (u *PaymentRepository) DeletePaymentByID(paymentID int64) error {
	return u.mysqlDb.Where("id = ?", paymentID).Delete(&model.Payment{}).Error
}

func (u *PaymentRepository) UpdatePayment(payment *model.Payment) error {
	return u.mysqlDb.Model(payment).Update(payment).Error
}

func (u *PaymentRepository) FindAll() (paymentAll []model.Payment, err error) {
	return paymentAll, u.mysqlDb.Find(&paymentAll).Error
}
