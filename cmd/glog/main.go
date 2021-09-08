package main

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/glog/cmd/glog/controller"
	"os"
)

func main() {
	e := echo.New()
	controller.InitRoute(e)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
