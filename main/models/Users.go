package models

import "fmt"

type Users struct {
	Id			int64
	Email 		string		`pg:",unique"`
	Password 	string
}

func (u Users) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}