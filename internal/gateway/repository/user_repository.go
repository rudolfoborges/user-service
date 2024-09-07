package repository

import (
	"context"

	"github.com/amandakeren/user-service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type postgresUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
        insert into users (id, name, email, password, active, created_at, updated_at)
        values (:id, :name, :email, :password, :active, :created_at, :updated_at)
    `

	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
        select * from users where email = $1 limit 1
    `

	var user entity.User

	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *postgresUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `
        select count(*) from users where email = $1
    `

	var count int64

	err := r.db.GetContext(ctx, &count, query, email)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
