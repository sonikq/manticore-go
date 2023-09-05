package db

import (
	"github.com/jmoiron/sqlx"
)

type AuthDB struct {
	*sqlx.DB
}

type IAuthDB interface {
	LoadAuthMatrix(id int) (*sqlx.Rows, error)
	CloseDB() error
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{
		DB: db,
	}
}

func (db *AuthDB) LoadAuthMatrix(id int) (*sqlx.Rows, error) {
	rows, err := db.Queryx(GetAuthMatrix)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (db *AuthDB) CloseDB() error {
	return db.Close()
}
