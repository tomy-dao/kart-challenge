package endpoint

import (
	"local/model"
	"local/service"
)

type ProductEndpoints interface {
	GetAllProducts() model.Response[[]*model.Product]
	GetProductById(id string) model.Response[*model.Product]
}

type productEndpoints struct {
	service service.Service
}	

func (e *productEndpoints) GetAllProducts() model.Response[[]*model.Product] {
	return e.service.Product.GetAllProducts()
}

func (e *productEndpoints) GetProductById(id string) model.Response[*model.Product] {
	return e.service.Product.GetProductById(id)
}

func NewProductEndpoints(params *Params) ProductEndpoints {
	return &productEndpoints{service: params.Service}
}
