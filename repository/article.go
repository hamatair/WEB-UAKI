package repository

import "github.com/gobuffalo/pop/v6"

type IArticleRepository interface {
	//isi daftar fungsi disini
	UploadArticle(title string, content string, authorID int, categoryID int) error
}

type ArticleRepository struct {
	db *pop.Connection
}

func NewArticleRepository(db *pop.Connection) IArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (r *ArticleRepository) UploadArticle(title string, content string, authorID int, categoryID int) error {
	// Implementasi fungsi untuk mengunggah artikel ke database

	return nil
}
