package controller

import (
	"encoding/json"
	"home/webonise/go/src/shopping_cart_app/structure"
	"html/template"
	"log"
	"net/http"
)

type UpdateShopProvider interface {
	UpdateToShop(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) UpdateToShop(w http.ResponseWriter, r *http.Request) {

	var p structure.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Service.UpdateShopProduct(p.ProdId, p.ProdPrice)
	if err != nil {
		log.Printf("failed to update product price : %v", err)
	}

	productitem, err := s.Service.RetriveShopProductByID(p.ProdId)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("view/updateshopproduct.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}

	t.Execute(w, productitem)
	w.WriteHeader(http.StatusOK)
}
