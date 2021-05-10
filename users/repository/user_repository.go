package repository

import (
	"context"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
)

type userRepository struct {
	db *config.DBConnection
}

func (u userRepository) Create(ctx context.Context, user *domain.Users) (int64, error) {
	db := u.db.Conn
	_, errorInsert := db.Model(user).Insert()
	return user.Id, errorInsert
}

func (u userRepository) Count(ctx context.Context, condition string, args ...interface{}) (int64, error) {
	var model = domain.Users{}
	return model.TotalRow(model, condition, args...)
}

func (u userRepository) First(context context.Context, conditions string, args ...interface{}) (*domain.Users, error) {
	users := &domain.Users{}
	db := u.db.Conn.Model(users)
	if conditions != ""{
		db.Where(conditions, args...)
	}
	err := db.First()
	return users, err
}

func (u userRepository) Select(context context.Context, conditions string, args ...interface{}) ([]domain.Users, error) {
	var users []domain.Users
	db := u.db.Conn.Model()
	if conditions != "" {
		db.Where(conditions, args...)
	}
	err := db.Select()
	return users, err
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{db: config.DBEngine}
}
