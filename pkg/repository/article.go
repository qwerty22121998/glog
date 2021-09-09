package repository

import (
	"context"
	"github.com/qwerty22121998/glog/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	oid, err := r.collection.InsertOne(ctx, article)
	if err != nil {
		return err
	}
	article.BaseModel.ID = oid.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *articleRepository) FindBySlug(ctx context.Context, slug string) (*model.Article, error) {
	var article *model.Article
	if err := r.collection.FindOne(ctx, bson.M{
		"slug": slug,
	}).Decode(&article); err != nil {
		return nil, err
	}
	return article, nil
}
