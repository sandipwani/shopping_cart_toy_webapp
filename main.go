package main

import (
	"net/http"

	"home/webonise/go/src/shopping_cart_app/Module"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/", Module.ProductInShop).Methods("GET")
	route.HandleFunc("/cart", Module.GetCartitems).Methods("GET")
	route.HandleFunc("/cart/add", Module.AddToCart).Methods("POST")
	route.HandleFunc("/cart/update", Module.UpdateCartitem).Methods("PUT")
	route.HandleFunc("/cart/delete", Module.DeleteCartitem).Methods("DELETE")

	http.ListenAndServe(":8888", route)
}
