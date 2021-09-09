package service

import (
	"context"
	"github.com/qwerty22121998/glog/pkg/dto"
	"github.com/qwerty22121998/glog/pkg/mapper"
	"github.com/qwerty22121998/glog/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ArticleService interface {
	Create(ctx context.Context, request dto.CreateArticleRequest) (*dto.Response, error)
}

type articleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(repoProvider repository.Provider) ArticleService {
	return &articleService{
		articleRepo: repoProvider.Article(),
	}
}

func (a *articleService) Create(ctx context.Context, request dto.CreateArticleRequest) (*dto.Response, error) {
	articleDTO := request.ToDTO()
	articleDTO.CreatedAt = time.Now()
	articleDTO.UpdatedAt = time.Now()
	_, err := a.articleRepo.FindBySlug(ctx, articleDTO.Slug)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, dto.InternalError(err)
	}
	if err == nil {
		return nil, dto.BadRequest("document with slug %v already existed", articleDTO.Slug)
	}

	articleModel := mapper.Article.ToModel(articleDTO)

	if err := a.articleRepo.Create(ctx, articleModel); err != nil {
		return nil, dto.InternalError(err)
	}

	return dto.Success(mapper.Article.ToDTO(articleModel)), nil
}
