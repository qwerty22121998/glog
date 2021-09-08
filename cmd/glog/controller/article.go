package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/glog/pkg/dto"
	"github.com/qwerty22121998/glog/pkg/service"
)

type ArticleController interface {
	Create(c echo.Context) error
}

type articleController struct {
	articleService service.ArticleService
}

func (a *articleController) Create(c echo.Context) error {

	ctx := c.Request().Context()
	var req dto.CreateArticleRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	resp := a.articleService.Create(ctx, req)

	return c.JSON(resp.Code(), resp.Response())
}
