package repository

import "github.com/gobuffalo/pop/v6"

type IArticleCategoryRepository interface {
	//isi daftar fungsi disini

}

type ArticleCategoryRepository struct {
	db *pop.Connection
}

func NewArticleCategoryRepository(db *pop.Connection) IArticleCategoryRepository {
	return &ArticleCategoryRepository{
		db: db,
	}
}
