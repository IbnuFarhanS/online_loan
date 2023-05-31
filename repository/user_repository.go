package repository

import "github.com/IbnuFarhanS/online_loan/model"

type UserRepository interface {
	Save(newUser model.User) (model.User, error)
	Update(updatedUser model.User) (model.User, error)
	Delete(id int64) (model.User, error)
	FindById(id int64) (model.User, error)
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
}
