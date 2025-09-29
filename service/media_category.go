package service

import "backend_server/repository"

type IMediaCategoryService interface {
	//isi daftar fungsi disini
}

type MediaCategoryService struct {
	mediacategoryrepository repository.IMediaCategoryRepository
}

func NewMediaCategoryService(mediacategoryrepository repository.IMediaCategoryRepository) IMediaCategoryService {
	return &MediaCategoryService{
		mediacategoryrepository: mediacategoryrepository,
	}
}
