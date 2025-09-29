package repository

import "github.com/gobuffalo/pop/v6"

type IAdminRepository interface {
	//isi daftar fungsi disini

}

type AdminRepository struct {
	db *pop.Connection
}

func NewAdminRepository(db *pop.Connection) IAdminRepository {
	return &AdminRepository{
		db: db,
	}
}
