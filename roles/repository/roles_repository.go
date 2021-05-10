package repository

import (
	"context"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
)

type repository struct {
	connection *config.DBConnection
}

func (r repository) BulkInsert(ctx context.Context, roles []*domain.Roles) error {
	for _, role := range roles {
		idInserted, errCreate := r.Store(ctx, role)
		if errCreate != nil{
			return errCreate
		}
		role.Id = idInserted
	}
	return nil
}

func (r repository) Store(ctx context.Context, roles *domain.Roles) (int64, error) {
	_, err := r.connection.Conn.Model(roles).Insert()
	return roles.Id, err
}

func NewRoleRepository() domain.RolesRepository {
	return &repository{connection: config.DBEngine}
}
