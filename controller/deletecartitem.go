package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type DeleteCartProvider interface {
	DeleteCartitem(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) DeleteCartitem(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	productcart, err := s.Service.RetriveCartProduct(id)
	if err != nil {
		log.Printf("failed to fetch product record from cart : %v", err)
	}

	err = s.Service.DeleteCartProduct(id)
	if err != nil {
		log.Printf("failed to delete product from cart %v", err)
	}

	t, err := template.ParseFiles("view/delete.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}
	t.Execute(w, productcart)

	w.WriteHeader(http.StatusOK)
}
