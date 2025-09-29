package repository

import "github.com/gobuffalo/pop/v6"

type IMediaCategoryRepository interface {
	//isi daftar fungsi disini

}

type MediaCategoryRepository struct {
	db *pop.Connection
}

func NewMediaCategoryRepository(db *pop.Connection) IMediaCategoryRepository {
	return &MediaCategoryRepository{
		db: db,
	}
}
