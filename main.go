package main

import (
	"fmt"
	"log"
	"pvg/controllers"
	"pvg/middlewares"
	"pvg/models"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	errLoadingEnvFile := godotenv.Load()
	if errLoadingEnvFile != nil {
		log.Fatal(errLoadingEnvFile)
	}

}
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic Detected : " + fmt.Sprint(err))
		}
	}()

	routing := echo.New()
	routing.Validator = &models.CustomValidator{Validator: validator.New()}

	route := routing.Group("/api/v1")

	controller := new(controllers.Controller)

	route.GET("/user/all", controller.GetAllUsers)
	route.GET("/user/:id", controller.GetUser)
	route.POST("/user", controller.CreateUser)
	route.PUT("/user/update/:id", controller.UpdateUser)
	route.DELETE("/user/:id", controller.DeleteUser)

	routing.HTTPErrorHandler = middlewares.RouteHandler
	routing.Logger.Fatal(routing.Start(":8000"))
}
