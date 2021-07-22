package main

import (
	"home/webonise/go/src/shopping_cart_app/containers"
	"home/webonise/go/src/shopping_cart_app/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "s"
	dbname   = "pgdata"
)

func main() {

	d := &database.PsqlIntializer{
		DBConfig: &database.DBConfig{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Dbname:   dbname,
		},
	}

	if err := d.InitializeConnection(); err != nil {
		log.Panic(err)
	}

	srv := &containers.Server{
		Router:   mux.NewRouter(),
		MasterDB: d.Db,
	}

	srv.InitializeServer()
	log.Printf("Starting server on the port 8888...")

	log.Fatal(http.ListenAndServe(":8888", srv.Router))
}
