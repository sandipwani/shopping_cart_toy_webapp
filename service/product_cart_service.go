package service

import (
	"home/webonise/go/src/shopping_cart_app/structure"
)

type ProductCartServiceProvider interface {
	RetriveCartProducts() (productlist []structure.ProductCart, err error)
	InsertCartProduct(productid int, productname string, productprice float64, productquantity int) error
	RetriveCartProduct(productid int) (product_info structure.ProductCart, err error)
	DeleteCartProduct(productid int) error
	UpdateCartProduct(productid int, productquantity int) error
	RetriveCartProductByID(productid int) (product_info structure.ProductCart, err error)
}

func (s *Service) RetriveCartProducts() (productlist []structure.ProductCart, err error) {
	return s.Accessor.RetriveCartProducts()
}

func (s *Service) InsertCartProduct(productid int, productname string, productprice float64, productquantity int) error {
	return s.Accessor.InsertCartProduct(productid, productname, productprice, productquantity)
}

func (s *Service) RetriveCartProduct(productid int) (product_info structure.ProductCart, err error) {
	return s.Accessor.RetriveCartProduct(productid)
}

func (s *Service) DeleteCartProduct(productid int) error {
	return s.Accessor.DeleteCartProduct(productid)
}

func (s *Service) UpdateCartProduct(productid int, productquantity int) error {
	return s.Accessor.UpdateCartProduct(productid, productquantity)
}

func (s *Service) RetriveCartProductByID(productid int) (product_info structure.ProductCart, err error) {
	return s.Accessor.RetriveCartProductByID(productid)
}
