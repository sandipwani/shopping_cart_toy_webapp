package Module

import (
	"fmt"
	"home/webonise/go/src/shopping_cart_app/Database"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func DeleteCartitem(w http.ResponseWriter, r *http.Request) {
	db := Database.OpenConnection()
	defer db.Close()

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var productcart Database.ProductCart
	err := db.QueryRow("SELECT *FROM product_cart WHERE product_id = $1;", id).Scan(&productcart.ProductId, &productcart.ProductName, &productcart.ProductPrice, &productcart.ProductQuantity)
	if err != nil {
		fmt.Printf("No corresponding record present in the product_cart table :: Re-run again & give different[present] id   ")
		log.Fatal(err)
	}
	sqlStatement := `DELETE FROM product_cart WHERE product_id = $1;`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	t, err := template.ParseFiles("Module/delete.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}
	t.Execute(w, productcart)

	w.WriteHeader(http.StatusOK)
}
