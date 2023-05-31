package usecase

import (
	"github.com/IbnuFarhanS/online_loan/model"
	"github.com/IbnuFarhanS/online_loan/repository"
)

type paymentMethodUsecase struct {
	paymentMethodRepository repository.PaymentMethodRepository
}

func NewPaymentMethodUsecase(paymentMethodRepository repository.PaymentMethodRepository) PaymentMethodUsecase {
	return &paymentMethodUsecase{
		paymentMethodRepository: paymentMethodRepository,
	}
}

func (u *paymentMethodUsecase) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	return u.paymentMethodRepository.Save(newPaymentMethod)
}

func (u *paymentMethodUsecase) Update(updatePaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	return u.paymentMethodRepository.Update(updatePaymentMethod)
}

func (u *paymentMethodUsecase) Delete(id int64) (model.PaymentMethod, error) {
	return u.paymentMethodRepository.Delete(id)
}

func (u *paymentMethodUsecase) FindById(id int64) (model.PaymentMethod, error) {
	return u.paymentMethodRepository.FindById(id)
}

func (u *paymentMethodUsecase) FindAll() ([]model.PaymentMethod, error) {
	return u.paymentMethodRepository.FindAll()
}

func (u *paymentMethodUsecase) FindByName(name string) (model.PaymentMethod, error) {
	return u.paymentMethodRepository.FindByName(name)
}
