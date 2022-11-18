package services

import (
	"errors"
	"math/rand"
	"net/http"
	"pvg/helpers"
	"pvg/models"
	"pvg/repository"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func NewServiceUser() repository.UserRepositories {
	return &UserService{}
}

func (u *UserService) GetAllUsers(ctx echo.Context) (*models.Responses, error) {
	user := new(models.User)
	result, err := user.GetAllUser()
	if err != nil {
		return &models.Responses{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "data found",
		Data:    result,
	}, nil

}

func (u *UserService) GetUser(ctx echo.Context) (*models.Responses, error) {
	user := new(models.User)

	user.Id = ctx.Param("id")
	result, err := user.GetOneUser()
	if err != nil {
		return &models.Responses{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "data found",
		Data:    result,
	}, nil
}

func (u *UserService) CreateUser(ctx echo.Context, request *models.UserCreate) (*models.Responses, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 0)
	if err != nil {
		return &models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	request.Password = string(password)

	result, err := request.CreateUser()
	if err != nil {
		return &models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	key := helpers.RandSeq(4)

	go func() {
		helpers.SendEmail("registration verification", request.Email, request.Username, key)
	}()

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "data found",
		Data:    result,
	}, nil
}

func (u *UserService) UpdateUser(ctx echo.Context, request *models.User) (*models.Responses, error) {

	request.Id = ctx.Param("id")

	result, err := request.UpdateUser()
	if err != nil {
		return &models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "data found",
		Data:    result,
	}, nil

}

func (u *UserService) DeleteUser(ctx echo.Context) (*models.Responses, error) {
	user := new(models.User)

	user.Id = ctx.Param("id")
	err := user.DeleteUser()
	if err != nil {
		return &models.Responses{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "data found",
	}, nil
}
