package repository

import (
	"gorm.io/gorm"

	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
)

type repository struct {
	connection *gorm.DB
}

func (r repository) Store(roles *domain.Roles) (uint, error) {
	result := r.connection.Create(roles)
	if result.Error != nil{
		return 0, result.Error
	}
	return roles.ID, nil
}

func NewRoleRepository() domain.RolesRepository {
	return &repository{connection: config.DBEngine}
}
