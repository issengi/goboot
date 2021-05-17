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
	var result int64
	err := u.db.QueryRow("INSERT INTO users(email, password) VALUES($1, $2) RETURNING id;",
		user.Email,
		user.Password).Scan(&result)
	return result, err
}

func (u userRepository) Count(condition string, args ...interface{}) (int64, error) {
	var model = &domain.Users{}
	return model.TotalRow(model, condition, args...)
}

func (u userRepository) First(conditions string, args ...interface{}) (*domain.Users, error) {
	users := &domain.Users{}
	query := fmt.Sprintf(`SELECT id, email, password, created_at, updated_at, deleted_at FROM users`)
	if conditions!=``{
		query = fmt.Sprintf(`%s WHERE %s`, query, conditions)
	}
	query += " LIMIT 1"
	err := u.db.QueryRow(query, args...).Scan(&users.Id,
		&users.Email,
		&users.Password,
		&users.CreatedAt,
		&users.UpdatedAt,
		&users.DeletedAt)
	return users, err
}

func (u userRepository) Select(conditions string, args ...interface{}) ([]domain.Users, error) {
	var users []domain.Users
	query := fmt.Sprintf(`SELECT id, email, created_at, updated_at, deleted_at FROM users`)
	if conditions != `` {
		query = fmt.Sprintf(`%s WHERE %s`, query, conditions)
	}
	rows, err := u.db.Query(query, args...)
	if err!=nil{
		return users, err
	}
	for rows.Next() {
		var item domain.Users
		errScan := rows.Scan(&item.Id, &item.Email, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if errScan != nil{
			return users, errScan
		}
		users = append(users, item)
	}

	return users, nil
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{db: config.DBEngine}
}
