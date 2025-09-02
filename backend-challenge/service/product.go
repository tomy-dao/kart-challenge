package service

import (
	"local/model"
	"local/repository"
)

type ProductService interface {
	GetAllProducts() model.Response[[]*model.Product]
	GetProductById(id string) model.Response[*model.Product]
}

type productService struct {
	repo repository.Repository
}

func (s *productService) GetAllProducts() model.Response[[]*model.Product] {
	return s.repo.Product().GetAllProducts()
}

func (s *productService) GetProductById(id string) model.Response[*model.Product] {
	return s.repo.Product().GetProductById(id)
}

func NewProductService(params *Params) ProductService {
	return &productService{repo: params.Repo}
}