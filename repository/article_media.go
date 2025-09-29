package repository

import "github.com/gobuffalo/pop/v6"

type IArticleMediaRepository interface {
	//isi daftar fungsi disini

}

type ArticleMediaRepository struct {
	db *pop.Connection
}

func NewArticleMediaRepository(db *pop.Connection) IArticleMediaRepository {
	return &ArticleMediaRepository{
		db: db,
	}
}
