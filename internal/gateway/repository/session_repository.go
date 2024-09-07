package repository

import (
	"context"

	"github.com/amandakeren/user-service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entity.Session) error
}

type postgresSessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) SessionRepository {
	return &postgresSessionRepository{db: db}
}

func (r *postgresSessionRepository) Create(ctx context.Context, session *entity.Session) error {
	query := `
        insert into sessions (user_id, created_at)
        values (:user_id, :created_at)
    `

	_, err := r.db.NamedExecContext(ctx, query, session)
	if err != nil {
		return err
	}

	return nil
}
