package repository

import "github.com/jmoiron/sqlx"

type PersonPostgres struct {
	db *sqlx.DB
}

func NewPersonPostgres(db *sqlx.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}


