package endpoint

import (
	"local/client"
	"local/service"
)


type Params struct {
	Service service.Service
	Client *client.Client
}

type Endpoints struct {
	Product ProductEndpoints
	Order OrderEndpoints
}

func NewEndpoints(params *Params) *Endpoints {
	return &Endpoints{
		Product: NewProductEndpoints(params),
		Order: NewOrderEndpoints(params),
	}
}
