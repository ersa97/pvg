package controllers

import (
	"net/http"
	"pvg/helpers"
	"pvg/models"
	"pvg/services"

	"github.com/labstack/echo/v4"
)

func (ct *Controller) GetAllUsers(ctx echo.Context) (err error) {

	service := services.NewServiceUser()
	data, err := service.GetAllUsers(ctx)

	if err != nil {
		return models.ResponseContext(data.Code, data.Message, nil, ctx)
	}

	return models.ResponseContext(data.Code, data.Message, data.Data, ctx)

}

func (ct *Controller) GetUser(ctx echo.Context) (err error) {

	service := services.NewServiceUser()
	data, err := service.GetUser(ctx)

	if err != nil {
		return models.ResponseContext(data.Code, data.Message, nil, ctx)
	}

	return models.ResponseContext(data.Code, data.Message, data.Data, ctx)

}

func (ct *Controller) CreateUser(ctx echo.Context) (err error) {

	user := new(models.UserCreate)
	if err = ctx.Bind(user); err != nil {
		helpers.LoggerError(err)
		return models.ResponseContext(http.StatusBadRequest, err.Error(), nil, ctx)
	}
	if err = ctx.Validate(user); err != nil {
		helpers.LoggerError(err)
		return models.ResponseContext(http.StatusBadRequest, err.Error(), nil, ctx)
	}

	service := services.NewServiceUser()
	data, err := service.CreateUser(ctx, user)

	if err != nil {
		return models.ResponseContext(data.Code, data.Message, nil, ctx)
	}
	return models.ResponseContext(data.Code, data.Message, data.Data, ctx)
}

func (ct *Controller) UpdateUser(ctx echo.Context) (err error) {

	user := new(models.User)

	user.Id = ctx.Param("id")

	if err = ctx.Bind(user); err != nil {
		helpers.LoggerError(err)
		return models.ResponseContext(http.StatusBadRequest, err.Error(), nil, ctx)
	}

	service := services.NewServiceUser()
	data, err := service.UpdateUser(ctx, user)

	if err != nil {
		return models.ResponseContext(data.Code, data.Message, nil, ctx)
	}
	return models.ResponseContext(data.Code, data.Message, data.Data, ctx)
}

func (ct *Controller) DeleteUser(ctx echo.Context) (err error) {

	service := services.NewServiceUser()
	data, err := service.DeleteUser(ctx)

	if err != nil {
		return models.ResponseContext(data.Code, data.Message, nil, ctx)
	}

	return models.ResponseContext(data.Code, data.Message, data.Data, ctx)

}
