package usecase

import (
	"github.com/IbnuFarhanS/online_loan/model"
	"github.com/IbnuFarhanS/online_loan/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) Save(newUser model.User) (model.User, error) {
	return u.userRepository.Save(newUser)
}

func (u *userUsecase) Update(updatedUser model.User) (model.User, error) {
	return u.userRepository.Update(updatedUser)
}

func (u *userUsecase) Delete(id int64) (model.User, error) {
	return u.userRepository.Delete(id)
}

func (u *userUsecase) FindById(id int64) (model.User, error) {
	return u.userRepository.FindById(id)
}

func (u *userUsecase) FindAll() ([]model.User, error) {
	return u.userRepository.FindAll()
}

func (u *userUsecase) FindByUsername(username string) (model.User, error) {
	return u.userRepository.FindByUsername(username)
}
