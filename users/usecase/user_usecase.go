package usecase

import (
	"context"
	"errors"
	"github.com/issengi/goboot/domain"
	userRoleRepository "github.com/issengi/goboot/user_role/repository"
	userRepository "github.com/issengi/goboot/users/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
	userRoleRepository domain.UserRoleRepository
}

func (u userUsecase) AssignRole(ctx context.Context, user *domain.Users, roles *domain.Roles) error {
	repo := u.userRoleRepository
	model := domain.UserRoles{RoleId: roles.Id, UserId: user.Id}
	return repo.Store(model)
}

func (u userUsecase) BulkInsert(ctx context.Context, users []*domain.Users) error {
	for _, user := range users {
		idInserted, errCreate := u.userRepository.Create(user)
		if errCreate!=nil{
			return errCreate
		}
		user.Id = idInserted
	}
	return nil
}

func (u userUsecase) Login(ctx context.Context, email, password string) (*domain.Users, error) {
	repository := u.userRepository
	user, errSelect := repository.First( `email = $1`, email)
	if errSelect != nil {
		return nil, errSelect
	}

	errorValidatePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errorValidatePassword != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func NewUserUsecase() domain.UserUsecase {
	u := userRepository.NewUserRepository()
	ur := userRoleRepository.NewUserRoleRepository()
	return &userUsecase{userRepository: u, userRoleRepository: ur}
}
