package repository

import (
	"context"
	"try/go-router/domain/entity"
)

type ArticleRepository interface {
	Store(ctx context.Context, dataArticle *entity.Article) error
	GetAllData(ctx context.Context, lang string, slug_category string) ([]*entity.Article, error)
	GetDataBySlug(ctx context.Context, codeArticle string) (*entity.Article, error)
}

type ArticleRedisRepository interface {
	GetAttributeArticleByKode(ctx context.Context, codeArticle string) (*entity.Article, error)
	StoreOrUpdateData(ctx context.Context, dataArticle *entity.Article) error
	GetAllData(ctx context.Context) (*entity.Article, error)
}
