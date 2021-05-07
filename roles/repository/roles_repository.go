package repository

import (
	"context"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
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
	var testId int64
	db := r.connection.Conn
	errorInsert := db.
		QueryRow(ctx, `INSERT INTO roles(role_name) VALUES($1) RETURNING id`, roles.Role).
		Scan(&testId)
	//defer db.Close(ctx)
	if errorInsert != nil{
		return 0, errorInsert
	}
	roles.Id = testId
	return roles.Id, nil
}

func NewRoleRepository() domain.RolesRepository {
	return &repository{connection: config.DBEngine}
}
