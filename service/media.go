package service

import "backend_server/repository"

type IMediaService interface {
	//isi daftar fungsi disini
}

type MediaService struct {
	mediarepository repository.IMediaRepository
}

func NewMediaService(mediarepository repository.IMediaRepository) IMediaService {
	return &MediaService{
		mediarepository: mediarepository,
	}
}
