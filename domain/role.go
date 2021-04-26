package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	Role string
	Users []*Users `gorm:"many2many:user_role;"`
}

type UserRole struct {
	RoleId int64
	UserId int64
}

func (r *Roles) String() string {
	return fmt.Sprintf("id %d, name %s", r.ID, r.Role)
}

func (r *Roles) GetName() string {
	return "Roles"
}

func (r *UserRole) GetName() string{
	return "User Role"
}

type RolesRepository interface {
	Store(r *Roles) (uint, error)
}
