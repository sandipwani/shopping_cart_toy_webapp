package service

import "home/webonise/go/src/shopping_cart_app/structure"

type ProductServiceProvider interface {
	InsertShopProduct(productid int, productname string, productprice float64) error
	RetriveShopProducts() (productlist []structure.Product, err error)
	RetriveShopProductByID(productid int) (product_info structure.ProductCart, err error)
}

func (s *Service) InsertShopProduct(productid int, productname string, productprice float64) error {
	return s.Accessor.InsertShopProduct(productid, productname, productprice)
}

func (s *Service) RetriveShopProducts() (productlist []structure.Product, err error) {
	return s.Accessor.RetriveShopProducts()
}

func (s *Service) RetriveShopProductByID(productid int) (product_info structure.ProductCart, err error) {
	return s.Accessor.RetriveShopProductByID(productid)
}
