package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
)

type repository struct {
	connection *pg.DB
}

func (r repository) Store(ctx context.Context, roles *domain.Roles) (int64, error) {
	_, err := r.connection.Model(roles).Insert()
	return roles.Id, err
}

func NewRoleRepository() domain.RolesRepository {
	return &repository{connection: config.DBEngine}
}
