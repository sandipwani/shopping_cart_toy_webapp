package controller

import (
	"encoding/json"
	"home/webonise/go/src/shopping_cart_app/structure"
	"html/template"
	"log"

	"net/http"
)

type AddCartProvider interface {
	AddToCart(w http.ResponseWriter, r *http.Request)
}

func (s *Srv) AddToCart(w http.ResponseWriter, r *http.Request) {

	var p structure.ProductCart
	//Try to decode the request body into the struct. If there is an error,
	//respond to the client with the error message and a 400 status code
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product_info, err := s.Service.RetriveShopProductByID(p.ProductId)
	if err != nil {
		log.Printf("failed to retrive product from shop : %v", err)
	}

	err = s.Service.InsertCartProduct(p.ProductId, product_info.ProductName, product_info.ProductPrice, p.ProductQuantity)
	if err != nil {
		log.Printf("failed to insert product into cart : %v", err)
	}

	productinfo := structure.ProductCart{ProductId: p.ProductId, ProductName: product_info.ProductName, ProductPrice: product_info.ProductPrice, ProductQuantity: p.ProductQuantity}
	t, err := template.ParseFiles("view/add.html")
	if err != nil {
		log.Fatal("html file reading err : ", err)
		return
	}

	t.Execute(w, productinfo)
	w.WriteHeader(http.StatusOK)
}
