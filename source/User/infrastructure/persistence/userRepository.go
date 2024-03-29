package persistence

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	domain "github.com/kudejwiktor/go-api-example/source/User/domain"
	"github.com/pkg/errors"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func (r UserRepositoryImpl) GetUserOfId(id int) (*domain.User, error) {
	var user domain.User
	if err := r.db.Get(&user, "SELECT * FROM users_db.users WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrEntityNotFound{Err: errors.WithStack(err)}
		}
		return nil, domain.ErrDbQuery{Err: errors.WithStack(err)}
	}
	return &user, nil
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &UserRepositoryImpl{db: db}
}
