package domain

import "fmt"

type Users struct {
	Id       int64
	Email    string `pg:",unique"`
	Password string
}

func (u Users) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}

type UserRepository interface {
	// First is select the first item where set condition
	First(conditions string, args ...interface{}) (*Users, error)
	// Select is list of user which descending ID
	Select(conditions string, args ...interface{}) ([]*Users, error)
}

type UserUsecase interface {
	Login(email, password string) (*Users, error)
}
