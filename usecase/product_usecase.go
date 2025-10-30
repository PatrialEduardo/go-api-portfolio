package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	//logica de negocio
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	//logica de negocio
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) GetProductByID(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(id_product)

	if err != nil {
		return nil, err
	}
	return product, nil
}