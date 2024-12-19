package repository

import (
	"context"
	"encelad-shared/domain"
	"encelad-shared/ports"
)

type UserPostgresRepositoryImpl struct {
	db ports.DBConn
}

func NewUserPostgresRepository(db ports.DBConn) *UserPostgresRepositoryImpl {
	return &UserPostgresRepositoryImpl{
		db: db,
	}
}

func (r *UserPostgresRepositoryImpl) rowToUserModel(row []any) *domain.UserModel {
	return domain.NewUserModel(
		row[0].(int64),
		row[1].(string),
		row[2].(string),
		domain.UserRole().FromStringMust(row[3].(string)),
	)
}

func (r *UserPostgresRepositoryImpl) GetByID(ctx context.Context, id int64) (*domain.UserModel, error) {
	query := "SELECT * FROM users WHERE id = $1"

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		row, err := rows.Values()
		if err != nil {
			return nil, err
		}
		return r.rowToUserModel(row), nil
	}

	return nil, ports.ErrNotFound
}

func (r *UserPostgresRepositoryImpl) Create(ctx context.Context, in *ports.CreateUserIn) (*domain.UserModel, error) {
	query := "INSERT INTO users (firstname, lastname) VALUES ($1, $2) RETURNING *"

	rows, err := r.db.Query(ctx, query, in.Firstname, in.Lastname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		row, err := rows.Values()
		if err != nil {
			return nil, err
		}
		return r.rowToUserModel(row), nil
	}

	return nil, ports.ErrNotFound
}

func (r *UserPostgresRepositoryImpl) Update(ctx context.Context, id int64, in *ports.UpdateUserIn) (*domain.UserModel, error) {
	query := "UPDATE users SET firstname=$1, lastname=$2 WHERE id=$3 RETURNING *"

	rows, err := r.db.Query(ctx, query, in.Firstname, in.Lastname, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		row, err := rows.Values()
		if err != nil {
			return nil, err
		}
		return r.rowToUserModel(row), nil
	}

	return nil, ports.ErrNotFound
}
