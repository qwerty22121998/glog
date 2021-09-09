package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/glog/pkg/dto"
	"github.com/qwerty22121998/glog/pkg/service"
	"net/http"
)

type ArticleController interface {
	Create(c echo.Context) error
}

type articleController struct {
	articleService service.ArticleService
}

func NewArticleController(serviceProvider service.Provider) ArticleController {
	return &articleController{
		articleService: serviceProvider.Article(),
	}
}

func (a *articleController) Create(c echo.Context) error {

	ctx := c.Request().Context()
	var req dto.CreateArticleRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	resp, err := a.articleService.Create(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
