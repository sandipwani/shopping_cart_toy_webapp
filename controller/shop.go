package controller

import (
	"html/template"
	"log"
	"net/http"
)

type ShopProductProvider interface {
	ProductInShop(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) ProductInShop(w http.ResponseWriter, r *http.Request) {

	productlist, err := s.Service.RetriveShopProducts()
	if err != nil {
		log.Printf("failed to fetch products from shop : %v", err)
	}

	t, err := template.ParseFiles("view/shop.html")
	if err != nil {
		log.Fatal("html file reading error : ", err)
		return
	}

	t.Execute(w, productlist)
}
