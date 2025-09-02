package client

import (
	"local/model"
)

type Client struct {}

func NewClient(params *model.InitParams) *Client {
	return &Client{}
}
