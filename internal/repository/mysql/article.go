package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/santekno/belajar-golang-restful/internal/models"
)

type ArticleStore struct {
	db *sql.DB
}

// New will create an object that represent the article_repository interface
func New(conn *sql.DB) *ArticleStore {
	return &ArticleStore{conn}
}

func (r *ArticleStore) GetAll(ctx context.Context) ([]*models.Article, error) {
	var result []*models.Article

	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		t := models.Article{}
		err = rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreateAt, &t.UpdateAt)

		if err != nil {
			return nil, err
		}

		result = append(result, &t)
	}

	return result, nil
}

func (r *ArticleStore) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	var result models.Article

	err := r.db.QueryRowContext(ctx, queryGetById, id).Scan(&result.ID, &result.Title, &result.Content, &result.CreateAt, &result.UpdateAt)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r *ArticleStore) Update(ctx context.Context, article *models.Article) (*models.Article, error) {
	res, err := r.db.ExecContext(ctx, queryUpdate, article.Title, article.Content, article.UpdateAt, article.ID)
	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	fmt.Printf("success update with affected %d", count)
	return article, nil
}

func (r *ArticleStore) Store(ctx context.Context, article *models.Article) (int64, error) {
	res, err := r.db.ExecContext(ctx, queryInsert, article.Title, article.Content, article.CreateAt, article.UpdateAt)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Printf("success create with lastId: %d", lastId)
	return lastId, nil
}

func (r *ArticleStore) Delete(ctx context.Context, id int64) (bool, error) {
	_, err := r.db.ExecContext(ctx, queryDelete, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
