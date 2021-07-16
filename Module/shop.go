package Module

import (
	"home/webonise/go/src/shopping_cart_app/Database"
	"html/template"
	"log"
	"net/http"
)

func ProductInShop(w http.ResponseWriter, r *http.Request) {
	db := Database.OpenConnection()
	rows, err := db.Query("SELECT *FROM product;")
	if err != nil {
		log.Fatal(err)
	}

	var productlist []Database.Product
	var productitem Database.Product
	for rows.Next() {

		rows.Scan(&productitem.ProdId, &productitem.ProdName, &productitem.ProdPrice)
		productlist = append(productlist, productitem)
	}

	t, err := template.ParseFiles("Module/shop.html")
	if err != nil {
		log.Fatal("html file reading error : ", err)
		return
	}
	t.Execute(w, productlist)

	defer rows.Close()
}
