package repository

import "github.com/IbnuFarhanS/online_loan/model"

type TransactionRepository interface {
	Save(newTransaction model.Transaction) (model.Transaction, error)
	Update(updateTransaction model.Transaction) (model.Transaction, error)
	Delete(id int64) (model.Transaction, error)
	FindById(id int64) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
	UpdateStatus(transactionID uint, status bool) error
	FindByUserID(userID int64) ([]model.Transaction, error)
}
