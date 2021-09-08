package repository

import (
	"context"
	"github.com/qwerty22121998/glog/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *model.Article) error
	FindBySlug(ctx context.Context, slug string) (*model.Article, error)
}
type articleRepository struct {
	collection *mongo.Collection
}

func NewArticleRepository(db *mongo.Database) ArticleRepository {
	return &articleRepository{
		collection: db.Collection("article"),
	}
}

func (r *articleRepository) Create(ctx context.Context, article *model.Article) error {
	_, err := r.collection.InsertOne(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) FindBySlug(ctx context.Context, slug string) (*model.Article, error) {
	article := &model.Article{
		Slug: slug,
	}
	if err := r.collection.FindOne(ctx, article).Decode(&article); err != nil {
		return nil, err
	}
	return article, nil
}
