package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/users/repository"
	"gitlab.com/NeoReids/backend-tryonline-golang/users/usecase"
	"time"
)

type Usecase interface {
	//Login(email, password string) (*domain.Users, error)
	CreateJWT(email, password string) (string, error)
}

type authRepository struct {
	con *pg.DB
}

func (a authRepository) CreateJWT(email, password string) (string, error) {
	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	user, errorFindUser := userUsecase.Login(email, password)
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

func NewAuthRepository(connection *pg.DB) Usecase {
	return &authRepository{con: connection}
}
