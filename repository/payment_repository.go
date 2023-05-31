package repository

import "github.com/IbnuFarhanS/online_loan/model"

type PaymentRepository interface {
	Save(newPayment model.Payment) (model.Payment, error)
	Update(updatePayment model.Payment) (model.Payment, error)
	Delete(id int64) (model.Payment, error)
	FindById(id int64) (model.Payment, error)
	FindAll() ([]model.Payment, error)
}
