package Module

import (
	"encoding/json"
	"home/webonise/go/src/shopping_cart_app/Database"
	"html/template"
	"log"
	"net/http"
)

func UpdateCartitem(w http.ResponseWriter, r *http.Request) {
	db := Database.OpenConnection()
	defer db.Close()

	var p Database.ProductCart
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `UPDATE product_cart SET product_quantity = $2 WHERE product_id=$1;`
	_, err = db.Exec(sqlStatement, p.ProductId, p.ProductQuantity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	var productcart Database.ProductCart
	err = db.QueryRow("SELECT *FROM product_cart WHERE product_id = $1;", p.ProductId).Scan(&productcart.ProductId, &productcart.ProductName, &productcart.ProductPrice, &productcart.ProductQuantity)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("Module/update.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}
	t.Execute(w, productcart)

	w.WriteHeader(http.StatusOK)
}
