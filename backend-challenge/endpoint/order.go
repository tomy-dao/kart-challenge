package endpoint

import (
	"local/model"
	"local/service"
	"strings"
)

type OrderEndpoints interface {
	PlaceAnOrder(req model.OrderRequest) model.Response[*model.Order]
}

type orderEndpoints struct {
	service service.Service
}

func (e *orderEndpoints) validateOrderRequest(req model.OrderRequest) model.Response[*model.Order] {
	if len(req.Items) == 0 {
		return model.Response[*model.Order]{
			Status: model.StatusBadRequest,
			Message: "at least one item is required",
		}
	}

	for _, item := range req.Items {
		if item.ProductID == "" {
			return model.Response[*model.Order]{
				Status: model.StatusBadRequest,
				Message: "productId for item is required",
			}
		}
		if item.Quantity <= 0 {
			return model.Response[*model.Order]{
				Status: model.StatusBadRequest,
				Message: "item quantity cannot be less than zero",
			}
		}
	}

	if req.CouponCode != "" {
		if len(req.CouponCode) < 8 || len(req.CouponCode) > 10 {
			return model.Response[*model.Order]{
				Status: model.StatusBadRequest,
				Message: "coupon code must be between 8 and 10 characters",
			}
		}
		validateCouponCodeRes := e.service.Order.ValidateCouponCode(strings.ToUpper(req.CouponCode))
		if validateCouponCodeRes.Error() {
			return model.Response[*model.Order]{
				Status: model.StatusBadRequest,
				Message: "invalid coupon code",
			}
		}
		if !validateCouponCodeRes.Data {
			return model.Response[*model.Order]{
				Status: model.StatusBadRequest,
				Message: "coupon code cannot be used",
			}
		}
	}

	return model.Response[*model.Order]{
		Status: model.StatusOK,
		Message: "valid order request",
	}
}

func (e *orderEndpoints) PlaceAnOrder(req model.OrderRequest) model.Response[*model.Order] {
	validateOrderRequestRes := e.validateOrderRequest(req)
	if validateOrderRequestRes.Error() {
		return validateOrderRequestRes
	}

	return e.service.Order.PlaceAnOrder(req)
}

func NewOrderEndpoints(params *Params) OrderEndpoints {
	return &orderEndpoints{service: params.Service}
}
