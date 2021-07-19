package Module

import (
	"encoding/json"
	"html/template"
	"log"

	"home/webonise/go/src/shopping_cart_app/Database"

	"net/http"
)

func AddToShop(w http.ResponseWriter, r *http.Request) {
	db := Database.OpenConnection()
	defer db.Close()

	var p Database.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO product (prod_id, prod_name, prod_price) VALUES ($1, $2, $3);`
	_, err = db.Exec(sqlStatement, p.ProdId, p.ProdName, p.ProdPrice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	productinfo := Database.Product{ProdId: p.ProdId, ProdName: p.ProdName, ProdPrice: p.ProdPrice}

	t, err := template.ParseFiles("Module/addshop.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}
	t.Execute(w, productinfo)

	w.WriteHeader(http.StatusOK)
}
