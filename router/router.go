package router

import (
	"home/webonise/go/src/shopping_cart_app/controller"

	"github.com/gorilla/mux"
)

type SrvRouter struct {
	Controller controller.SrvController
	Router     *mux.Router
}

func (r *SrvRouter) InitilizeRoutes() {

	r.Router.HandleFunc("/shop", r.Controller.ShowShopProducts).Methods("GET")
	r.Router.HandleFunc("/shop/add", r.Controller.AddToShop).Methods("POST")
	r.Router.HandleFunc("/shop/update", r.Controller.UpdateToShop).Methods("PUT")
	r.Router.HandleFunc("/shop/delete", r.Controller.DeleteFromShop).Methods("DELETE")
	
	r.Router.HandleFunc("/cart", r.Controller.ShowCartProducts).Methods("GET")
	r.Router.HandleFunc("/cart/add", r.Controller.AddToCart).Methods("POST")
	r.Router.HandleFunc("/cart/update", r.Controller.UpdateToCart).Methods("PUT")
	r.Router.HandleFunc("/cart/delete", r.Controller.DeleteFromCart).Methods("DELETE")
}
