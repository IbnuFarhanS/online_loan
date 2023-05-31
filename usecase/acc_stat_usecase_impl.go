package usecase

import (
	"github.com/IbnuFarhanS/online_loan/model"
	"github.com/IbnuFarhanS/online_loan/repository"
)

type acceptStatusUsecase struct {
	acceptStatusRepository repository.AcceptStatusRepository
}

func NewAcceptStatusUsecase(acceptStatusRepository repository.AcceptStatusRepository) AcceptStatusUsecase {
	return &acceptStatusUsecase{
		acceptStatusRepository: acceptStatusRepository,
	}
}

func (u *acceptStatusUsecase) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	return u.acceptStatusRepository.Save(newAcceptStatus)
}

func (u *acceptStatusUsecase) Update(updateAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	return u.acceptStatusRepository.Update(updateAcceptStatus)
}

func (u *acceptStatusUsecase) Delete(id int64) (model.AcceptStatus, error) {
	return u.acceptStatusRepository.Delete(id)
}

func (u *acceptStatusUsecase) FindById(id int64) (model.AcceptStatus, error) {
	return u.acceptStatusRepository.FindById(id)
}

func (u *acceptStatusUsecase) FindAll() ([]model.AcceptStatus, error) {
	return u.acceptStatusRepository.FindAll()
}
