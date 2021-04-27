package usecase

import (
	"context"
	"errors"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
	"gitlab.com/NeoReids/backend-tryonline-golang/users/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func (u userUsecase) BulkInsert(ctx context.Context, users []domain.Users) error {
	for _, user := range users {
		_, errCreate := u.userRepository.Create(ctx, &user)
		if errCreate!=nil{
			return errCreate
		}
	}
	return nil
}

func (u userUsecase) Login(ctx context.Context, email, password string) (*domain.Users, error) {
	repository := u.userRepository
	user, errSelect := repository.First(ctx, "email = ?", email)
	if errSelect != nil {
		return nil, errors.New("failed to get record from that identification")
	}

	errorValidatePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errorValidatePassword != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func NewUserUsecase() domain.UserUsecase {
	u := repository.NewUserRepository()
	return &userUsecase{userRepository: u}
}
