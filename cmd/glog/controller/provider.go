package controller

import "github.com/qwerty22121998/glog/pkg/service"

type Provider interface {
	Article() ArticleController
}

type provider struct {
	articleController ArticleController
}

func NewProvider(serviceProvider service.Provider) Provider {
	return &provider{
		articleController: NewArticleController(serviceProvider),
	}
}

func (p *provider) Article() ArticleController {
	return p.articleController
}
