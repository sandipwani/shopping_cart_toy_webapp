package Module

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"home/webonise/go/src/shopping_cart_app/Database"

	"net/http"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {

	db := Database.OpenConnection()
	defer db.Close()

	//declare a new ProductCart struct
	var p Database.ProductCart
	//Try to decode the request body into the struct. If there is an error,
	//respond to the client with the error message and a 400 status code
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if p.ProductId < 101 && p.ProductId > 110 {
		fmt.Println("Please give the ProductId in range between 101-110 ***")
	}

	var product_info Database.ProductCart
	err = db.QueryRow("SELECT *FROM product WHERE prod_id = $1;", p.ProductId).Scan(&product_info.ProductId, &product_info.ProductName, &product_info.ProductPrice)
	if err != nil {
		log.Fatal("Failed to execute select query to get product from seller: ", err)
	}

	sqlStatement := `INSERT INTO product_cart (product_id, product_name, product_price, product_quantity) VALUES ($1, $2, $3, $4);`
	_, err = db.Exec(sqlStatement, product_info.ProductId, product_info.ProductName, product_info.ProductPrice, p.ProductQuantity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	productinfo := Database.ProductCart{ProductId: p.ProductId, ProductName: product_info.ProductName, ProductPrice: product_info.ProductPrice, ProductQuantity: p.ProductQuantity}

	t, err := template.ParseFiles("Module/add.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}
	t.Execute(w, productinfo)

	w.WriteHeader(http.StatusOK)
}
