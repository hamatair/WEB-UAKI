package service

import "backend_server/repository"

type IArticleMediaService interface {
	//isi daftar fungsi disini
}

type ArticleMediaService struct {
	articlemediarepository repository.IArticleMediaRepository
}

func NewArticleMediaService(articlemediarepository repository.IArticleMediaRepository) IArticleMediaService {
	return &ArticleMediaService{
		articlemediarepository: articlemediarepository,
	}
}
