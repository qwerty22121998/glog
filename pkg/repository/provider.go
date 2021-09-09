package repository

import "go.mongodb.org/mongo-driver/mongo"

type Provider interface {
	Article() ArticleRepository
}

type provider struct {
	articleRepository ArticleRepository
}

func NewProvider(db *mongo.Database) Provider {
	return &provider{
		articleRepository: NewArticleRepository(db),
	}
}

func (p *provider) Article() ArticleRepository {
	return p.articleRepository
}
