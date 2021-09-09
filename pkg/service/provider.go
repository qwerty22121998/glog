package service

import "github.com/qwerty22121998/glog/pkg/repository"

type Provider interface {
	Article() ArticleService
}

type provider struct {
	articleService ArticleService
}

func NewProvider(repoProvider repository.Provider) Provider {
	return &provider{
		articleService: NewArticleService(repoProvider),
	}
}

func (p *provider) Article() ArticleService {
	return p.articleService
}
