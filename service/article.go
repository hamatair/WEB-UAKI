package service

import "backend_server/repository"

type IArticleService interface {
	//isi daftar fungsi disini
}

type ArticleService struct {
	articlerepository repository.IArticleRepository
}

func NewArticleService(articlerepository repository.IArticleRepository) IArticleService {
	return &ArticleService{
		articlerepository: articlerepository,
	}
}
