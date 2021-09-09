package dto

import "github.com/qwerty22121998/glog/pkg/util"

type Article struct {
	BaseDTO
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Content     string   `json:"content"`
}

type CreateArticleRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Content     string   `json:"content" validate:"required"`
}

func (r *CreateArticleRequest) ToDTO() *Article {
	return &Article{
		Slug:        util.GenerateSlug(r.Title),
		Title:       r.Title,
		Description: r.Description,
		Tags:        r.Tags,
		Content:     r.Content,
	}
}
