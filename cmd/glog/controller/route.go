package controller

import "github.com/labstack/echo/v4"

func InitRoute(e *echo.Echo) {
	var article ArticleController

	e.POST("/article", article.Create)
}
