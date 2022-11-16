package controllers

import (
	"pvg/repository"
	"pvg/services"
)

type Controller struct {
	RepoUser repository.UserRepositories
}

func NewController() *Controller {
	repo := services.NewServiceUser()

	return &Controller{
		RepoUser: repo,
	}
}
