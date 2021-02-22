package article

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/author"
	"github.com/bxcodec/go-clean-arch/domain"
)

type service struct {
	repo           Repository
	authorRepo     author.Repository
	contextTimeout time.Duration
}

// NewService will create new an service object representation of Service interface
func NewService(repo Repository, authorRepo author.Repository, timeout time.Duration) Service {
	return service{
		repo:           repo,
		authorRepo:     authorRepo,
		contextTimeout: timeout,
	}
}

// Count ...
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query ...
func (s service) Query(ctx context.Context, offset, limit int) ([]domain.Article, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	result := []domain.Article{}
	for _, item := range items {
		author, _ := s.authorRepo.GetByID(ctx, item.Author.ID)
		item.Author = author
		result = append(result, item)
	}

	return result, err
}
