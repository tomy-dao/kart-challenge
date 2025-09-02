package model


type Order struct {
	ID string `json:"id"`
	CouponCode string `json:"couponCode"`
	Items []*Item `json:"items"`
	Products []*Product `json:"products"`
}

type Item struct {
	ProductID string `json:"productId"`
	Quantity int `json:"quantity"`
}

type OrderRequest struct {
	CouponCode string `json:"couponCode"`
	Items []*Item `json:"items"`
}
