package repository

import (
	"pvg/models"

	"github.com/labstack/echo/v4"
)

type UserRepositories interface {
	GetAllUsers(ctx echo.Context) (*models.Responses, error)
	GetUser(ctx echo.Context) (*models.Responses, error)
	CreateUser(ctx echo.Context, request *models.UserCreate) (*models.Responses, error)
	UpdateUser(ctx echo.Context, request *models.User) (*models.Responses, error)
	DeleteUser(ctx echo.Context) (*models.Responses, error)
}
