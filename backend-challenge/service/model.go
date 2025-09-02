package service

import (
	"local/client"
	"local/repository"
)

type Params struct {
	Repo   repository.Repository
	Client *client.Client
}