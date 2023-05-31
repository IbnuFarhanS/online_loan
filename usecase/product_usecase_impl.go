package usecase

import (
	"github.com/IbnuFarhanS/online_loan/model"
	"github.com/IbnuFarhanS/online_loan/repository"
)

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (u *productUsecase) Save(newProduct model.Product) (model.Product, error) {
	return u.productRepository.Save(newProduct)
}

func (u *productUsecase) Update(updateProduct model.Product) (model.Product, error) {
	return u.productRepository.Update(updateProduct)
}

func (u *productUsecase) Delete(id int64) (model.Product, error) {
	return u.productRepository.Delete(id)
}

func (u *productUsecase) FindById(id int64) (model.Product, error) {
	return u.productRepository.FindById(id)
}

func (u *productUsecase) FindAll() ([]model.Product, error) {
	return u.productRepository.FindAll()
}

func (u *productUsecase) FindByName(name string) (model.Product, error) {
	return u.productRepository.FindByName(name)
}
