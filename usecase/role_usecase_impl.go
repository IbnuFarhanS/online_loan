package usecase

import (
	"github.com/IbnuFarhanS/online_loan/model"
	"github.com/IbnuFarhanS/online_loan/repository"
)

type roleUsecase struct {
	roleRepository repository.RoleRepository
}

func NewRoleUsecase(roleRepo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepository: roleRepo,
	}
}

func (u *roleUsecase) Save(newRole model.Role) (model.Role, error) {
	return u.roleRepository.Save(newRole)
}

func (u *roleUsecase) Update(updatedUser model.Role) (model.Role, error) {
	return u.roleRepository.Update(updatedUser)
}

func (u *roleUsecase) Delete(id int64) (model.Role, error) {
	return u.roleRepository.Delete(id)
}

func (u *roleUsecase) FindById(id int64) (model.Role, error) {
	return u.roleRepository.FindById(id)
}

func (u *roleUsecase) FindAll() ([]model.Role, error) {
	return u.roleRepository.FindAll()
}

func (u *roleUsecase) FindByName(name string) (model.Role, error) {
	return u.roleRepository.FindByName(name)
}
