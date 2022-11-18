package services

import (
	"encoding/base64"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"pvg/helpers"
	"pvg/models"
	"pvg/repository"
	"strings"
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

	emailkey := helpers.CreateKey(request.Email, key)

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "data found",
		Data: map[string]interface{}{
			"EmailKey": emailkey,
			"User":     result,
		},
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

func (u *UserService) ConfirmEmail(ctx echo.Context, request *models.ConfirmEmail) (*models.Responses, error) {

	key, err := base64.StdEncoding.DecodeString(request.EmailKey)
	emailkey := strings.Split(string(key), "|")

	if err != nil {
		return &models.Responses{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	log.Println(emailkey)
	if emailkey[1] != request.Key {
		return &models.Responses{
			Code:    http.StatusBadRequest,
			Message: "email key not match",
			Data:    nil,
		}, errors.New("email key not match")
	}

	user := new(models.User)

	user.Email = emailkey[0]
	user.EmailConfirmed = 1

	err = user.ConfirmEmail()
	if err != nil {
		return &models.Responses{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "email confirmed",
	}, nil
}

func (u *UserService) ForgotInitial(ctx echo.Context, request *models.ForgotInitial) *models.Responses {

	rand.Seed(time.Now().UnixNano())
	key := helpers.RandSeq(4)

	forgotkey := helpers.CreateKey(request.Email, key)

	go func() {
		helpers.SendEmail("forgot password verification", request.Email, request.Username, key)
	}()

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "email key sent",
		Data:    forgotkey,
	}
}

func (u *UserService) ForgotValidation(ctx echo.Context, request *models.ForgotValidation) (*models.Responses, error) {

	key, err := base64.StdEncoding.DecodeString(request.ForgotKey)
	emailkey := strings.Split(string(key), "|")

	if err != nil {
		return &models.Responses{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	log.Println(emailkey)
	if emailkey[1] != request.Key {
		return &models.Responses{
			Code:    http.StatusBadRequest,
			Message: "email key not match",
			Data:    nil,
		}, errors.New("email key not match")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 0)
	if err != nil {
		return &models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	user := new(models.User)

	user.Password = string(password)

	user.Email = emailkey[0]
	user.EmailConfirmed = 1

	err = user.EditPassword()
	if err != nil {
		return &models.Responses{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}, errors.New(err.Error())
	}

	return &models.Responses{
		Code:    http.StatusOK,
		Message: "password changed",
	}, nil
}
