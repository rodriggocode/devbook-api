package database

import (
	"database/sql"
	"devbook-api/app/config"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", config.StringConnectDatabase)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
