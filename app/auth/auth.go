package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/issengi/goboot/app/config"
	userRoleRepository "github.com/issengi/goboot/user_role/repository"
	"github.com/issengi/goboot/users/usecase"
	"time"
)

type ClaimersObject struct {
	UserId   	int64 		`json:"user_id"`
	Remember 	bool 		`json:"remember"`
	Role		[]string 	`json:"role"`
	jwt.StandardClaims
}

type Usecase interface {
	//Login(email, password string) (*domain.Users, error)
	CreateJWT(ctx context.Context, email, password string) (string, error)
}

type authRepository struct {
	con *config.DBConnection
}

func (a authRepository) CreateJWT(ctx context.Context, email, password string) (string, error) {
	userUsecase := usecase.NewUserUsecase()
	userRole := userRoleRepository.NewUserRoleRepository()
	user, errorFindUser := userUsecase.Login(ctx, email, password)
	if errorFindUser != nil {
		return "", errorFindUser
	}
	roles, errFindRole := userRole.SelectReturnRole(`users_id = $1`, user.Id)
	if errFindRole!=nil{
		return "", errFindRole
	}
	var listRole []string
	for _, role := range roles {
		listRole = append(listRole, role.Role)
	}
	claimers := ClaimersObject{
		UserId:   user.Id,
		Remember: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Id:        uuid.NewString(),
			Issuer:    config.Config.AppName,
		},
		Role: listRole,
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
