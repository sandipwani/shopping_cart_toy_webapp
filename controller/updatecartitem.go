package controller

import (
	"encoding/json"
	"home/webonise/go/src/shopping_cart_app/structure"
	"html/template"
	"log"
	"net/http"
)

type UpdateCartProvider interface {
	UpdateCartitem(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) UpdateCartitem(w http.ResponseWriter, r *http.Request) {

	var p structure.ProductCart
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Service.UpdateCartProduct(p.ProductId, p.ProductQuantity)
	if err != nil {
		log.Printf("failed to update product : %v", err)
	}

	productitem, err := s.Service.RetriveCartProductByID(p.ProductId)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("view/update.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}

	t.Execute(w, productitem)
	w.WriteHeader(http.StatusOK)
}
