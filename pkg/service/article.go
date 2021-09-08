package service

import (
	"context"
	"github.com/qwerty22121998/glog/pkg/dto"
	"github.com/qwerty22121998/glog/pkg/mapper"
	"github.com/qwerty22121998/glog/pkg/repository"
	"time"
)

type ArticleService interface {
	Create(ctx context.Context, request dto.CreateArticleRequest) dto.CommonResponse
}

type articleService struct {
	articleRepo repository.ArticleRepository
}

func (a *articleService) Create(ctx context.Context, request dto.CreateArticleRequest) dto.CommonResponse {
	articleDTO := request.ToDTO()
	articleDTO.CreatedAt = time.Now()
	articleDTO.UpdatedAt = time.Now()

	articleModel := mapper.Article.ToModel(articleDTO)

	if err := a.articleRepo.Create(ctx, articleModel); err != nil {
		return dto.InternalError(err)
	}

	return dto.Success(mapper.Article.ToDTO(articleModel))
}
