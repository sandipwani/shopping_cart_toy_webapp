package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PsqlIntializer struct {
	Db *sql.DB
	*DBConfig
}

func (i *PsqlIntializer) InitializeConnection() error {
	i.Driver = "postgres"
	db, err := i.DBConfig.connect(i.connstr())
	if err != nil {
		log.Printf("failed to connect db : %v", err)
		return err
	}

	i.Db = db
	return nil
}

func (i *PsqlIntializer) connstr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ", i.Host, i.Port, i.User, i.Password, i.Dbname)
}
