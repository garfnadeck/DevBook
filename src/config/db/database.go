package db

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		err := db.Close()
		return nil, err
	}
	return db, nil
}
