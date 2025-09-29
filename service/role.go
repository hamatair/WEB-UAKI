package service

import "backend_server/repository"

type IRoleService interface {
	//isi daftar fungsi disini
}

type RoleService struct {
	rolerepository repository.IRoleRepository
}

func NewRoleService(rolerepository repository.IRoleRepository) IRoleService {
	return &RoleService{
		rolerepository: rolerepository,
	}
}
