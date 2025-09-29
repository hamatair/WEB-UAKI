package repository

import "github.com/gobuffalo/pop/v6"

type IStatusRepository interface {
	//isi daftar fungsi disini

}

type StatusRepository struct {
	db *pop.Connection
}

func NewStatusRepository(db *pop.Connection) IStatusRepository {
	return &StatusRepository{
		db: db,
	}
}
