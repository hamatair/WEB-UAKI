package service

import "backend_server/repository"

type IAdminService interface {
	//isi daftar fungsi disini
}

type AdminService struct {
	adminrepository repository.IAdminRepository
}

func NewAdminService(adminrepository repository.IAdminRepository) IAdminService {
	return &AdminService{
		adminrepository: adminrepository,
	}
}
