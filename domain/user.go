package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id       int64
	Email    string `gorm:"unique"`
	Password string
	Role   	[]*Roles `gorm:"many2many:user_role;"`
}

func (u Users) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}

func (u Users) GetName() string{
	return fmt.Sprintf("Users")
}

type UserRepository interface {
	// First is select the first item where set condition
	First(conditions string, args ...interface{}) (*Users, error)
	// Select is list of user which descending ID
	Select(conditions string, args ...interface{}) ([]Users, error)
}

type UserUsecase interface {
	Login(email, password string) (*Users, error)
}
