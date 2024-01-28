package repository

import "github.com/jmoiron/sqlx"

type Person interface {
	
}

type Repository struct {
	Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person: NewPersonPostgres(db),
	}
}