package controller

import (
	"html/template"
	"log"
	"net/http"
)

type GetCartProvider interface {
	GetCartitems(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) GetCartitems(w http.ResponseWriter, r *http.Request) {

	productlist, err := s.Service.RetriveCartProducts()
	if err != nil {
		log.Printf("failed to fetch products from shop : %v", err)
		return
	}

	t, err := template.ParseFiles("view/cart.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}
	t.Execute(w, productlist)
}
