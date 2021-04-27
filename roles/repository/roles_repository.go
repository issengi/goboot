package repository

import (
	"context"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
)

type repository struct {
	connection *config.DBConnection
}

func (r repository) BulkInsert(ctx context.Context, roles []domain.Roles) error {
	for _, role := range roles {
		_, errCreate := r.Store(ctx, &role)
		if errCreate != nil{
			return errCreate
		}
	}
	return nil
}

func (r repository) Store(ctx context.Context, roles *domain.Roles) (int64, error) {
	db := r.connection.Conn
	errorInsert := db.QueryRow(ctx, `INSERT INTO roles(role) VALUES($1); RETURNING id`, roles.Role).Scan(roles.Id)
	defer db.Close(ctx)
	if errorInsert != nil{
		return 0, errorInsert
	}
	return roles.Id, nil
}

func NewRoleRepository() domain.RolesRepository {
	return &repository{connection: config.DBEngine}
}
