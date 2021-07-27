package controller

import "home/webonise/go/src/shopping_cart_app/service"

type Srv struct {
	Service service.ServiceProvider
}

type SrvController interface {
	AddCartProvider
	AddShopProvider
	UpdateShopProvider
	DeleteShopProvider
	DeleteCartProvider
	GetCartProvider
	ShopProductProvider
	UpdateCartProvider
}
