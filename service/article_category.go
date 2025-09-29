package service

import "backend_server/repository"

type IArticleCategoryService interface {
	//isi daftar fungsi disini
}

type ArticleCategoryService struct {
	articlecategoryrepository repository.IArticleCategoryRepository
}

func NewArticleCategoryService(articlecategoryrepository repository.IArticleCategoryRepository) IArticleCategoryService {
	return &ArticleCategoryService{
		articlecategoryrepository: articlecategoryrepository,
	}
}
