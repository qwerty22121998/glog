package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/glog/cmd/glog/api"
	"github.com/qwerty22121998/glog/cmd/glog/provider"
	"github.com/qwerty22121998/glog/pkg/client"
	"log"
	"net/http"
	"os"
)

type customValidator struct {
	validator *validator.Validate
}

func (c *customValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := client.NewMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	providers := provider.NewProvider(db)

	e := echo.New()

	e.Validator = &customValidator{validator: validator.New()}

	api.InitRoute(e, providers.Controller())
	address := fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(address))
}
