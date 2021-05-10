package domain

import (
	"context"
	"fmt"
)

type Users struct {
	Id       	int64
	Email    	string `gorm:"unique"`
	Password 	string
	BaseModel
}

func (u Users) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}

func (u Users) GetName() string{
	return fmt.Sprintf("users")
}

type UserRepository interface {
	// First is select the first item where set condition
	First(context context.Context, conditions string, args ...interface{}) (*Users, error)
	// Select is list of user which descending ID
	Select(context context.Context, conditions string, args ...interface{}) ([]Users, error)
	// Count all user match with condition
	Count(ctx context.Context, condition string, args ...interface{}) (int64, error)
	// Create new user
	Create(ctx context.Context, user *Users) (int64, error)
}

type UserUsecase interface {
	Login(ctx context.Context, email, password string) (*Users, error)
	BulkInsert(ctx context.Context, users []*Users) error
	AssignRole(ctx context.Context, user *Users, roles *Roles) error
}
