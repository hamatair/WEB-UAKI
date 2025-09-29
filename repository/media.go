package repository

import "github.com/gobuffalo/pop/v6"

type IMediaRepository interface {
	//isi daftar fungsi disini
}

type MediaRepository struct {
	db *pop.Connection
}

func NewMediaRepository(db *pop.Connection) IMediaRepository {
	return &MediaRepository{
		db: db,
	}
}
