package repository

import (
	"context"

	"github.com/santekno/learn-golang-restful/internal/models"
)

//go:generate mockery --name=ArticleRepository --filename=repository_mock.go --inpackage

type ArticleRepository interface {
	GetAll(ctx context.Context) ([]*models.Article, error)
	GetByID(ctx context.Context, id int64) (*models.Article, error)
	Update(ctx context.Context, article *models.Article) (*models.Article, error)
	Store(ctx context.Context, article *models.Article) (int64, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
