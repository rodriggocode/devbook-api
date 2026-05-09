package database

import (
	"database/sql"
	"devbook-api/app/config"
	"time"

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

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(5 * time.Minute)

	return conn, nil
}
