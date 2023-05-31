package usecase

import "github.com/IbnuFarhanS/online_loan/model"

type ProductUsecase interface {
	Save(newProduct model.Product) (model.Product, error)
	Update(updateProduct model.Product) (model.Product, error)
	Delete(id int64) (model.Product, error)
	FindById(id int64) (model.Product, error)
	FindAll() ([]model.Product, error)
	FindByName(name string) (model.Product, error)
}
