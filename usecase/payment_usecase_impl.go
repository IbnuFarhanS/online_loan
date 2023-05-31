package usecase

import (
	"github.com/IbnuFarhanS/online_loan/model"
	"github.com/IbnuFarhanS/online_loan/repository"
)

type paymentUsecase struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentUsecase(paymentRepository repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}

func (u *paymentUsecase) Save(newPayment model.Payment) (model.Payment, error) {
	return u.paymentRepository.Save(newPayment)
}

func (u *paymentUsecase) Update(updatePayment model.Payment) (model.Payment, error) {
	return u.paymentRepository.Update(updatePayment)
}

func (u *paymentUsecase) Delete(id int64) (model.Payment, error) {
	return u.paymentRepository.Delete(id)
}

func (u *paymentUsecase) FindById(id int64) (model.Payment, error) {
	return u.paymentRepository.FindById(id)
}

func (u *paymentUsecase) FindAll() ([]model.Payment, error) {
	return u.paymentRepository.FindAll()
}
