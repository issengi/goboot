package cmd

import (
	"context"
	"github.com/issengi/goboot/domain"
	roleRepository "github.com/issengi/goboot/roles/repository"
	"github.com/issengi/goboot/users/usecase"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	userUsecase := usecase.NewUserUsecase()
	roleRepo := roleRepository.NewRoleRepository()
	ctx, cancelContext := context.WithCancel(context.Background())
	defer cancelContext()
	var roles = []domain.Roles{
		{Role: "admin"},
		{Role: "user"},
	}
	errSeedRole := roleRepo.BulkInsert(ctx, roles)
	if errSeedRole != nil{
		panic(errSeedRole)
	}
	passwordAdmin, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	passwordUser, _ := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	var users = []domain.Users{
		{Email: "admin@example.com", Password: string(passwordAdmin)},
		{Email: "user@example.com", Password: string(passwordUser)},
	}

	errSeedUser := userUsecase.BulkInsert(ctx, users)
	if errSeedUser != nil{
		panic(errSeedUser)
	}
}