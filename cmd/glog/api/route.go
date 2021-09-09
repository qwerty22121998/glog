package api

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/glog/cmd/glog/controller"
	"net/http"
)

func InitRoute(e *echo.Echo, controllers controller.Provider) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/article/:slug", controllers.Article().FindBySlug)
	e.POST("/article", controllers.Article().Create)
}
