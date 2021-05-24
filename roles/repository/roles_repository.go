package repository

import (
	"fmt"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
)

type repository struct {
	connection *config.DBConnection
}

func (r repository) Select(where string, args ...interface{}) ([]domain.Roles, error) {
	query := fmt.Sprintf(`SELECT * FROM roles`)
	if where != `` {
		query = fmt.Sprintf(`%s WHERE %s`, query, where)
	}

	rows, err := r.connection.Conn.Queryx(query, args...)
	defer rows.Close()

	if err!=nil{
		return nil, err
	}
	var roles []domain.Roles
	for rows.Next() {
		var item domain.Roles
		errScan := rows.Scan(&item.Id, &item.Role)
		if errScan != nil{
			return nil, errScan
		}
		roles = append(roles, item)
	}
	return roles, nil
}

func (r repository) BulkInsert(roles []*domain.Roles) error {
	for _, role := range roles {
		idInserted, errCreate := r.Store(role)
		if errCreate != nil{
			return errCreate
		}
		role.Id = idInserted
	}
	return nil
}

func (r repository) Store(roles *domain.Roles) (int64, error) {
	return r.connection.InsertReturnId(`INSERT INTO roles(role_name) VALUES(:role) RETURNING id`, map[string]interface{}{
		"role": roles.Role,
	})
}

func NewRoleRepository() domain.RolesRepository {
	return &repository{connection: config.DBEngine}
}
