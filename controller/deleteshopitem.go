package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type DeleteShopProvider interface {
	DeleteFromShop(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) DeleteFromShop(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	productcart, err := s.Service.RetriveShopProductByID(id)
	if err != nil {
		log.Printf("failed to fetch product record from shop : %v", err)
	}

	err = s.Service.DeleteShopProduct(id)
	if err != nil {
		log.Printf("failed to delete product from shop %v", err)
	}

	t, err := template.ParseFiles("view/deleteshopitem.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}

	t.Execute(w, productcart)
	w.WriteHeader(http.StatusOK)
}
