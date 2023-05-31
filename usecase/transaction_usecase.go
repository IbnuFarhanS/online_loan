package usecase

import "github.com/IbnuFarhanS/online_loan/model"

type TransactionUsecase interface {
	Save(newTransaction model.Transaction, userid int64) (model.Transaction, error)
	Update(updateTransaction model.Transaction) (model.Transaction, error)
	Delete(id int64) (model.Transaction, error)
	FindById(id int64) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
}
