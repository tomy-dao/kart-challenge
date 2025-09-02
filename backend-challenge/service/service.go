package service

type Service struct {
	Product ProductService
	Order OrderService
}


func NewService(params *Params) Service {
	return Service{
		Product: NewProductService(params),
		Order: NewOrderService(params),
	}
}
