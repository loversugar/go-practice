package service

import (
	"go-practice/electricity-project/datamodels"
	"go-practice/electricity-project/repositories"
)

type IProductService interface {
	GetProductById(int64) (*datamodels.Product, error)

	GetAllProducts() ([]*datamodels.Product, error)

	DeleteProductById(int64) bool

	InsertProduct(product *datamodels.Product)(int64,error)

	UpdateProduct(product *datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

func NewProductServcice(repo repositories.IProduct) IProductService {
	return &ProductService{productRepository:repo}
}

func (p *ProductService) GetProductById(productId int64)(*datamodels.Product, error) {
	return p.productRepository.SelectByKey(productId)
}

func (p *ProductService) GetAllProducts() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}

func (p *ProductService) DeleteProductById(productId int64) bool  {
	return p.productRepository.Delete(productId)
}

func (p *ProductService) InsertProduct(product *datamodels.Product)(int64,error) {
	return p.productRepository.Insert(product)
}

func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}
