package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Driver   string
}

type DBConnectionInitializer interface {
	InitializeConnection() error
}

func (db *DBConfig) connect(connStr string) (*sql.DB, error) {
	return sql.Open(db.Driver, connStr)
}
