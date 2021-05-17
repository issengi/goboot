package domain

import (
	"fmt"
	"github.com/go-pg/pg/v10/orm"
)

func init(){
	orm.RegisterTable((*UserRoles)(nil))
}

type Roles struct {
	BaseModel
	Id		int64
	Role 	string
	Users 	[]*Users `gorm:"many2many:user_roles;"`
}

type UserRoles struct {
	RoleId int64	`pg:"roles_id"`
	UserId int64	`pg:"users_id"`
}

func (r *Roles) String() string {
	return fmt.Sprintf("id %d, name %s", r.Id, r.Role)
}

func (r *Roles) GetName() string {
	return "roles"
}

func (r *UserRoles) GetName() string{
	return "user_roles"
}

type RolesRepository interface {
	Store(r *Roles) (int64, error)
	BulkInsert(roles []*Roles) error
	Select(where string, args ...interface{})([]Roles, error)
}

type UserRoleRepository interface {
	Store(userRoleStruct UserRoles) error
	SelectReturnRole(condition string, args ...interface{})([]Roles, error)
	SelectReturnUser(condition string, args ...interface{})([]Users, error)
}