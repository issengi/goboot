package cmd

import (
	"context"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
	roleRepository "gitlab.com/NeoReids/backend-tryonline-golang/roles/repository"
	"gitlab.com/NeoReids/backend-tryonline-golang/users/usecase"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
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
		log.Printf(errSeedRole.Error())
		os.Exit(1)
	}
	passwordAdmin, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	passwordUser, _ := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	var users = []domain.Users{
		{Email: "admin@example.com", Password: string(passwordAdmin)},
		{Email: "user@example.com", Password: string(passwordUser)},
	}

	errSeedUser := userUsecase.BulkInsert(ctx, users)
	if errSeedUser != nil{
		log.Printf(errSeedUser.Error())
		os.Exit(1)
	}
}