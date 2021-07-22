package controller

import (
	"encoding/json"
	"html/template"
	"log"

	"home/webonise/go/src/shopping_cart_app/structure"

	"net/http"
)

type AddShopProvider interface {
	AddToShop(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) AddToShop(w http.ResponseWriter, r *http.Request) {

	var p structure.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Service.InsertShopProduct(p.ProdId, p.ProdName, p.ProdPrice)
	if err != nil {
		log.Printf("failed to insert product into shop table : %v", err)
	}

	productinfo := structure.Product{ProdId: p.ProdId, ProdName: p.ProdName, ProdPrice: p.ProdPrice}

	t, err := template.ParseFiles("view/addshop.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}

	t.Execute(w, productinfo)
	w.WriteHeader(http.StatusOK)
}
