package repository

import "github.com/gobuffalo/pop/v6"

type IRoleRepository interface {
	//isi daftar fungsi disini

}

type RoleRepository struct {
	db *pop.Connection
}

func NewRoleRepository(db *pop.Connection) IRoleRepository {
	return &RoleRepository{
		db: db,
	}
}
