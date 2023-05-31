package usecase

import (
	"github.com/IbnuFarhanS/online_loan/model"
	"github.com/IbnuFarhanS/online_loan/repository"
)

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
	}
}

func (u *transactionUsecase) Save(newTransaction model.Transaction, userID int64) (model.Transaction, error) {
	return u.transactionRepository.Save(newTransaction)
}

func (u *transactionUsecase) Update(updateTransaction model.Transaction) (model.Transaction, error) {
	return u.transactionRepository.Update(updateTransaction)
}

func (u *transactionUsecase) Delete(id int64) (model.Transaction, error) {
	return u.transactionRepository.Delete(id)
}

func (u *transactionUsecase) FindById(id int64) (model.Transaction, error) {
	return u.transactionRepository.FindById(id)
}

func (u *transactionUsecase) FindAll() ([]model.Transaction, error) {
	return u.transactionRepository.FindAll()
}
