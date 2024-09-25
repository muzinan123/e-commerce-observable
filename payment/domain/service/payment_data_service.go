package service

import (
	"git.imooc.com/coding-447/payment/domain/model"
	"git.imooc.com/coding-447/payment/domain/repository"
)

type IPaymentDataService interface {
	AddPayment(*model.Payment) (int64, error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
	FindAllPayment() ([]model.Payment, error)
}

func NewPaymentDataService(paymentRepository repository.IPaymentRepository) IPaymentDataService {
	return &PaymentDataService{paymentRepository}
}

type PaymentDataService struct {
	PaymentRepository repository.IPaymentRepository
}

func (u *PaymentDataService) AddPayment(payment *model.Payment) (int64, error) {
	return u.PaymentRepository.CreatePayment(payment)
}

func (u *PaymentDataService) DeletePayment(paymentID int64) error {
	return u.PaymentRepository.DeletePaymentByID(paymentID)
}

func (u *PaymentDataService) UpdatePayment(payment *model.Payment) error {
	return u.PaymentRepository.UpdatePayment(payment)
}

func (u *PaymentDataService) FindPaymentByID(paymentID int64) (*model.Payment, error) {
	return u.PaymentRepository.FindPaymentByID(paymentID)
}

func (u *PaymentDataService) FindAllPayment() ([]model.Payment, error) {
	return u.PaymentRepository.FindAll()
}
