package author

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

// Repository represent ontract
type Repository interface {
	GetByID(ctx context.Context, id int64) (domain.Author, error)
}
