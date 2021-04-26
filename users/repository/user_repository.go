package repository

import (
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) First(conditions string, args ...interface{}) (*domain.Users, error) {
	list, err := u.Select(conditions, args...)
	if err != nil {
		return nil, err
	}
	return &list[0], err
}

func (u userRepository) Select(conditions string, args ...interface{}) ([]domain.Users, error) {
	var users []domain.Users
	db := u.db
	tx := db.Where(conditions, args...).Find(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{db: config.DBEngine}
}
