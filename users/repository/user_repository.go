package repository

import (
	"fmt"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
)

type userRepository struct {
	db *config.DBConnection
}

func (u userRepository) Create(user *domain.Users) (int64, error) {
	return u.db.InsertReturnId(`INSERT INTO users(email, password) VALUES(:email, :password) RETURNING id`,
		map[string]interface{}{
			"email": user.Email,
			"password": user.Password,
		})
}

func (u userRepository) Count(condition string, args ...interface{}) (int64, error) {
	var model = &domain.Users{}
	return model.TotalRow(model, condition, args...)
}

func (u userRepository) First(conditions string, args ...interface{}) (*domain.Users, error) {
	var user domain.Users
	query := fmt.Sprintf(`SELECT id, email, password, created_at, updated_at, deleted_at FROM users`)
	if conditions!=``{
		query = fmt.Sprintf(`%s WHERE %s`, query, conditions)
	}
	query += " LIMIT 1"
	err := u.db.Conn.Get(&user, query, args...)
	return &user, err
}

func (u userRepository) Select(conditions string, args ...interface{}) ([]domain.Users, error) {
	users := []domain.Users{}
	query := fmt.Sprintf(`SELECT id, email, created_at, updated_at, deleted_at FROM users`)
	if conditions != `` {
		query = fmt.Sprintf(`%s WHERE %s`, query, conditions)
	}
	err := u.db.Conn.Get(&users, query, args...)
	return users, err
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{db: config.DBEngine}
}
