package domain

import (
	"context"
	"fmt"
)

type Roles struct {
	BaseModel
	Id		int64
	Role 	string
	Users 	[]*Users `gorm:"many2many:user_role;"`
}

type UserRole struct {
	RoleId int64
	UserId int64
}

func (r *Roles) String() string {
	return fmt.Sprintf("id %d, name %s", r.Id, r.Role)
}

func (r *Roles) GetName() string {
	return "roles"
}

func (r *UserRole) GetName() string{
	return "User Role"
}

type RolesRepository interface {
	Store(ctx context.Context, r *Roles) (int64, error)
	BulkInsert(ctx context.Context, roles []Roles) error
}
