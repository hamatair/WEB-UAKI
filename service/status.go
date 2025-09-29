package service

import "backend_server/repository"

type IStatusService interface {
	//isi daftar fungsi disini
}

type StatusService struct {
	statusrepository repository.IStatusRepository
}

func NewStatusService(statusrepository repository.IStatusRepository) IStatusService {
	return &StatusService{
		statusrepository: statusrepository,
	}
}
