package service

import "backend_server/repository"

type Service struct {
	AdminService           IAdminService
	RoleService            IRoleService
	StatusService          IStatusService
	ArticleService         IArticleService
	ArticleCategoryService IMediaCategoryService
	ArticleMediaService    IArticleMediaService
	MediaService           IMediaService
	MediaCategoryService   IMediaCategoryService
}

type InitParam struct {
	repository *repository.Repository
}

func NewService(param InitParam) *Service {
	return &Service{
		AdminService:           NewAdminService(param.repository.AdminRepository),
		RoleService:            NewRoleService(param.repository.RoleRepository),
		StatusService:          NewStatusService(param.repository.StatusRepository),
		ArticleService:         NewArticleService(param.repository.ArticleRepository),
		ArticleCategoryService: NewMediaCategoryService(param.repository.MediaCategoryRepository),
		ArticleMediaService:    NewArticleMediaService(param.repository.ArticleMediaRepository),
		MediaService:           NewMediaService(param.repository.MediaRepository),
		MediaCategoryService:   NewMediaCategoryService(param.repository.MediaCategoryRepository),
	}
}
