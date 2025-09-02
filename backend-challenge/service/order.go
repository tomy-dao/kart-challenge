package service

import (
	"fmt"
	"local/model"
	"local/repository"
	initdata "local/repository/data"

	"github.com/google/uuid"
)

type OrderService interface {
	PlaceAnOrder(req model.OrderRequest) model.Response[*model.Order]
	ValidateCouponCode(code string) model.Response[bool]
}

type orderService struct {
	repo repository.Repository
}

func (s *orderService) PlaceAnOrder(req model.OrderRequest) model.Response[*model.Order] {
	order := &model.Order{
		ID: uuid.New().String(),
		CouponCode: req.CouponCode,
		Items: req.Items,
	}

	products := make([]*model.Product, 0)
	for _, item := range req.Items {
		prd := s.repo.Product().GetProductById(item.ProductID)
		fmt.Println(prd)
		if prd.Error() {
			return model.Response[*model.Order]{
				Data: nil,
				Status: model.StatusUnprocessableEntity,
				Message: "invalid product specified",
			}
		}
		products = append(products, prd.Data)
	}

	order.Products = products

	return model.SuccessResponse[*model.Order](order)
}

func (s *orderService) ValidateCouponCode(code string) model.Response[bool] {
	valid := initdata.CheckCodeExistsInAllowedCodes(code)
	return model.Response[bool]{
		Data: valid,
		Status: model.StatusOK,
		Message: "Coupon code is valid",
	}
}


func NewOrderService(params *Params) OrderService {
	return &orderService{repo: params.Repo}
}