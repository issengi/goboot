package repository

import (
	"context"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
)

type userRoleRepository struct {
	db *config.DBConnection
}

func (u userRoleRepository) Store(ctx context.Context, userRoleStruct domain.UserRole) error {
	db := u.db.Conn
	_, err := db.Exec(ctx,
		`INSERT INTO user_role(role_id, user_id) VALUES($1, $2)`,
		userRoleStruct.RoleId, userRoleStruct.UserId)
	return err
}

func NewUserRoleRepository() domain.UserRoleRepository {
	return &userRoleRepository{db: config.DBEngine}
}