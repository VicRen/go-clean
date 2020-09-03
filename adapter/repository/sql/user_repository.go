package sql

import (
	"context"
	"database/sql"

	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/repository"
)

func NewUserRepository(conn *sql.DB) repository.UserRepository {
	return &userRepository{Conn: conn}
}

type userRepository struct {
	Conn *sql.DB
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	query := `SELECT id,name,email FROM user`
	users, err := r.fetch(context.Background(), query, nil)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []entity.User, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result = make([]entity.User, 0)
	for rows.Next() {
		u := entity.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}
