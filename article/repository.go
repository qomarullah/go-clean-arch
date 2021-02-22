package article

import (
	"context"
	"database/sql"
	"log"

	"github.com/bxcodec/go-clean-arch/domain"
)

type repository struct {
	DB *sql.DB
}

// NewRepository will create an object that represent the Repository interface
func NewRepository(conn *sql.DB) Repository {
	return repository{conn}
}

// Count ..
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	sql := "SELECT count(*) from article"
	err := r.DB.QueryRow(sql).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, err
}

// Query ...
func (r repository) Query(ctx context.Context, offset, limit int) ([]domain.Article, error) {
	var items []domain.Article

	sql := "SELECT id, title, author_id, content, updated_at, created_at FROM article  limit ? offset ?"
	log.Println("sql", sql, limit, offset)
	results, err := r.DB.Query(sql, limit, offset)
	if err != nil {
		return items, err
	}
	for results.Next() {
		var item domain.Article
		err = results.Scan(&item.ID, &item.Title, &item.Author.ID, &item.Content, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}

	return items, err
}
