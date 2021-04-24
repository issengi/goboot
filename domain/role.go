package domain

import (
	"context"
	"fmt"
)

type Roles struct {
	Id   int64
	Role string
}

type UserRole struct {
	RoleId int64
	UserId int64
	Role   *Roles `pg:"rel"`
}

func (r *Roles) String() string {
	return fmt.Sprintf("id %d, name %s", r.Id, r.Role)
}

type RolesRepository interface {
	Store(ctx context.Context, r *Roles) (int64, error)
}
