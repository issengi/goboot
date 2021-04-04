package entities

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/issengi/goboot/main/config"
	"github.com/issengi/goboot/main/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthRepository interface {
	Login(email, password string) (*models.Users, error)
	CreateJWT(users *models.Users) (string, error)
}

type authRepository struct {
	con *pg.DB
}

func (a authRepository) Login(email, password string) (*models.Users, error) {
	user := new(models.Users)
	err := a.con.Model(user).Where("email = ?", email).Select()
	if err!=nil{
		return user, errors.New("User not found")
	}

	errorValidatePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errorValidatePassword!=nil{
		return user, errors.New("Wrong password")
	}

	return user, nil
}

func (a authRepository) CreateJWT(users *models.Users) (string, error) {
	claimers := struct {
		UserId int64
		Remember bool
		jwt.StandardClaims
	}{
		UserId: users.Id,
		Remember: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Id: uuid.NewString(),
			Issuer: config.Config.AppName,
		},
	}
	sign := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claimers,
	)
	token, errSigned := sign.SignedString([]byte(config.Config.AppKey))
	if errSigned!=nil{
		return "", errSigned
	}

	return token, nil

}

func NewAuthRepository(connection *pg.DB) AuthRepository {
	return &authRepository{con: connection}
}