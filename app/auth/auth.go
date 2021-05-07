package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/users/usecase"
	"time"
)

type Usecase interface {
	//Login(email, password string) (*domain.Users, error)
	CreateJWT(ctx context.Context, email, password string) (string, error)
}

type authRepository struct {
	con *config.DBConnection
}

func (a authRepository) CreateJWT(ctx context.Context, email, password string) (string, error) {
	userUsecase := usecase.NewUserUsecase()
	user, errorFindUser := userUsecase.Login(ctx, email, password)
	if errorFindUser != nil {
		return "", errorFindUser
	}

	claimers := struct {
		UserId   int64
		Remember bool
		jwt.StandardClaims
	}{
		UserId:   user.Id,
		Remember: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Id:        uuid.NewString(),
			Issuer:    config.Config.AppName,
		},
	}
	sign := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claimers,
	)
	token, errSigned := sign.SignedString([]byte(config.Config.AppKey))
	if errSigned != nil {
		return "", errSigned
	}

	return token, nil

}

func NewAuthRepository() Usecase {
	connection := config.DBEngine
	return &authRepository{con: connection}
}
