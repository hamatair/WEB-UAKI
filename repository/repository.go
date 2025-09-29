package repository

import "github.com/gobuffalo/pop/v6"

type Repository struct {
	AdminRepository   IAdminRepository
	RoleRepository    IRoleRepository
	ArticleRepository IArticleRepository
	ArticleCategoryRepository IArticleCategoryRepository
	StatusRepository  IStatusRepository
	ArticleMediaRepository IArticleMediaRepository
	MediaRepository   IMediaRepository
	MediaCategoryRepository   IMediaCategoryRepository
}

func NewRepository(db *pop.Connection) *Repository {
	return &Repository{
		AdminRepository:   NewAdminRepository(db),
		RoleRepository:    NewRoleRepository(db),
		ArticleRepository: NewArticleRepository(db),
		ArticleCategoryRepository: NewArticleCategoryRepository(db),
		StatusRepository:  NewStatusRepository(db),
		ArticleMediaRepository: NewArticleMediaRepository(db),
		MediaRepository:   NewMediaRepository(db),
		MediaCategoryRepository:   NewMediaCategoryRepository(db),
	}
}
