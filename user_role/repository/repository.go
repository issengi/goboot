package repository

import (
	"fmt"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/domain"
	rolerepo "github.com/issengi/goboot/roles/repository"
	usersrepo "github.com/issengi/goboot/users/repository"
)

type userRoleRepository struct {
	db *config.DBConnection
	roleRepository domain.RolesRepository
	userRepository domain.UserRepository
}

func (u userRoleRepository) SelectReturnRole(condition string, args ...interface{}) ([]domain.Roles, error) {
	subquery := fmt.Sprintf(`SELECT roles_id FROM user_role`)
	if condition!=`` {
		subquery = fmt.Sprintf(`%s WHERE %s`, subquery, condition)
	}
	return u.roleRepository.Select(fmt.Sprintf(`id IN (%s)`, subquery), args...)
}

func (u userRoleRepository) SelectReturnUser(condition string, args ...interface{}) ([]domain.Users, error) {
	subquery := fmt.Sprintf(`SELECT users_id FROM user_role`)
	if condition!=`` {
		subquery = fmt.Sprintf(`%s WHERE %s`, subquery, condition)
	}
	return u.userRepository.Select(fmt.Sprintf(`id IN (%s)`, subquery), args)
}

func (u userRoleRepository) Store(userRoleStruct domain.UserRoles) error {
	_, err := u.db.Conn.NamedExec(`INSERT INTO user_roles(roles_id, users_id) VALUES(:roleid, :userid)`, map[string]interface{}{
		"roleid": userRoleStruct.RoleId,
		"userid": userRoleStruct.UserId,
	})
	return err
}

func NewUserRoleRepository() domain.UserRoleRepository {
	roleRepo := rolerepo.NewRoleRepository()
	userRepo := usersrepo.NewUserRepository()
	return &userRoleRepository{db: config.DBEngine, userRepository: userRepo, roleRepository: roleRepo}
}
