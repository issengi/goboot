package repository

import (
	"github.com/go-pg/pg/v10"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
)

type userRepository struct {
	db *pg.DB
}

func (u userRepository) First(conditions string, args ...interface{}) (*domain.Users, error) {
	list, err := u.Select(conditions, args...)
	if err != nil {
		return nil, err
	}
	return list[0], err
}

func (u userRepository) Select(conditions string, args ...interface{}) ([]*domain.Users, error) {
	var users []*domain.Users
	db := u.db
	defer db.Close()
	errorSelect := db.Model(&users).Where(conditions, args...).Select()
	if errorSelect != nil {
		return users, errorSelect
	}
	return users, nil
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{db: config.DBEngine}
}
