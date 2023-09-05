package db

import (
	"fmt"

	"gitlab.geogracom.com/skdf/skdf-manticore-go/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driver_name = "postgres"
)

type DB struct {
	IAuthDB
	IRoadDB
}

func NewDB(config configs.Config) (*DB, error) {
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.DBName, config.DB.SSLMode)

	db, err := sqlx.Open(driver_name, connInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{
		IAuthDB: NewAuthDB(db),
		IRoadDB: NewRoadDB(db),
	}, nil
}

func (db *DB) Close() error {
	if err := db.IAuthDB.CloseDB(); err != nil {
		return err
	}

	if err := db.IRoadDB.CloseDB(); err != nil {
		return err
	}

	return nil
}
