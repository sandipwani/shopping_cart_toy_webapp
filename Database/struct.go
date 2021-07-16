package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Product struct {
	ProdId    int     `json:"prod_id"`
	ProdName  string  `json:"prod_name"`
	ProdPrice float64 `json:"prod_price"`
}

type ProductCart struct {
	ProductId       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float64 `json:"product_price"`
	ProductQuantity int     `json:"product_quantity"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "s"
	dbname   = "pgdata"
)

func OpenConnection() *sql.DB {

	psqlInfo := fmt.Sprintf("host= %s port=%d user=%s password =%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//fmt.Println("Succesfully connected to database!")
	return db
}
