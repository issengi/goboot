package usecase

import (
	"errors"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func (u userUsecase) Login(email, password string) (*domain.Users, error) {
	repository := u.userRepository
	user, errSelect := repository.First("email = ?", email)
	if errSelect != nil {
		return nil, errors.New("failed to get record from that identification")
	}

	errorValidatePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errorValidatePassword != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepository: u}
}
