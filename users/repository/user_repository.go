package repository

import (
	"context"
	"fmt"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
	"github.com/jackc/pgx/v4"
)

type userRepository struct {
	db *config.DBConnection
}

func (u userRepository) Create(ctx context.Context, user *domain.Users) (int64, error) {
	query := "INSERT INTO users(email, password) VALUES($1,$2) RETURNING id"
	var lastInserted int64
	errCreate := u.db.Conn.QueryRow(ctx, query, user.Email, user.Password).Scan(&lastInserted)
	return lastInserted, errCreate
}

func (u userRepository) Count(ctx context.Context, condition string, args ...interface{}) (int64, error) {
	var model = domain.Users{}
	return model.TotalRow(model, condition, args...)
}

func (u userRepository) First(context context.Context, conditions string, args ...interface{}) (*domain.Users, error) {
	var query = fmt.Sprintf("SELECT id, email, password, created_at, updated_at, deleted_at FROM users %s", conditions)
	if conditions == ""{
		query = fmt.Sprintf("SELECT id, email, password, created_at, updated_at, deleted_at FROM users ORDER BY id DESC")
	}
	user := &domain.Users{}
	err := u.db.Conn.QueryRow(context, query, args...).Scan(user.Id, user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRepository) Select(context context.Context, conditions string, args ...interface{}) ([]*domain.Users, error) {
	db := u.db.Conn
	var users []*domain.Users
	var query = "SELECT id, email, password, created_at, updated_at, deleted_at FROM users %s ORDER BY id DESC"
	var errorQuery error
	var rows pgx.Rows
	defer rows.Close()
	if conditions == "" {
		rows, errorQuery = db.Query(context, query)
	}else{
		query = fmt.Sprintf(query, conditions)
		rows, errorQuery = db.Query(context, query, args...)
	}
	if errorQuery!=nil{
		return nil, errorQuery
	}
	for rows.Next() {
		user := &domain.Users{}
		errorScan := rows.Scan(user.Id, user.Email,
			user.Password,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt)
		if errorScan!=nil{
			return nil, errorScan
		}
		users = append(users, user)
	}
	return users, nil
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{db: config.DBEngine}
}
