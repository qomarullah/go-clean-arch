package author

import (
	"context"
	"database/sql"
	"log"

	"github.com/bxcodec/go-clean-arch/domain"
)

type repository struct {
	DB *sql.DB
}

// NewRepository will create an implementation of Repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		DB: db,
	}
}

// GetByID ...
func (r *repository) GetByID(ctx context.Context, id int64) (domain.Author, error) {

	var author domain.Author
	sql := "SELECT * FROM author where id=?"
	log.Println("sqlAuthor", sql, id)
	err := r.DB.QueryRow(sql, id).Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt)

	return author, err
}
