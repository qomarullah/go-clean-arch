package article

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

// Usecase ...
type Service interface {
	Query(ctx context.Context, offset, limit int) ([]domain.Article, error)
	Count(ctx context.Context) (int, error)
	//GetByID(ctx context.Context, id int64) (domain.Article, error)
	//Update(ctx context.Context, ar *domain.Article) error
	//GetByTitle(ctx context.Context, title string) (domain.Article, error)
	//Store(context.Context, *domain.Article) error
	//Delete(ctx context.Context, id int64) error
}

// Repository ...
type Repository interface {
	// Count returns the number of customers.
	Count(ctx context.Context) (int, error)
	// Query returns the list of customers with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]domain.Article, error)
	//Fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Article, err error)
	//GetByID(ctx context.Context, id int64) (domain.Article, error)
	//GetByTitle(ctx context.Context, title string) (domain.Article, error)
	//Update(ctx context.Context, ar *domain.Article) error
	//Store(ctx context.Context, a *domain.Article) error
	//Delete(ctx context.Context, id int64) error
}
