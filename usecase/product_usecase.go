package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	//Repository
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo, 
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	//GetProducts
	return pu.repository.GetProducts()
}