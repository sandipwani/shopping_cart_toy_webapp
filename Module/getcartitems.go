package Module

import (
	"home/webonise/go/src/shopping_cart_app/Database"
	"html/template"
	"log"
	"net/http"
)

func GetCartitems(w http.ResponseWriter, r *http.Request) {
	db := Database.OpenConnection()
	rows, err := db.Query("SELECT *FROM product_cart;")
	if err != nil {
		log.Fatal(err)
	}

	var productitem []Database.ProductCart
	var productcart Database.ProductCart
	for rows.Next() {

		rows.Scan(&productcart.ProductId, &productcart.ProductName, &productcart.ProductPrice, &productcart.ProductQuantity)
		productitem = append(productitem, productcart)
	}

	//templating
	//PageTitle := "****This is the list of products from cart****"
	t, err := template.ParseFiles("Module/cart.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}
	t.Execute(w, productitem)

	defer rows.Close()
}
