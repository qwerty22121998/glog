package mapper

import (
	"github.com/qwerty22121998/glog/pkg/dto"
	"github.com/qwerty22121998/glog/pkg/model"
)

type ArticleMapper int

var Article ArticleMapper

func (ArticleMapper) ToDTO(article *model.Article) *dto.Article {
	if article == nil {
		return nil
	}
	return &dto.Article{
		BaseDTO:     Base.ToDTO(article.BaseModel),
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		Tags:        article.Tags,
		Content:     article.Content,
	}
}

func (ArticleMapper) ToModel(article *dto.Article) *model.Article {
	if article == nil {
		return nil
	}
	return &model.Article{
		BaseModel:   Base.ToModel(article.BaseDTO),
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		Tags:        article.Tags,
		Content:     article.Content,
	}
}
