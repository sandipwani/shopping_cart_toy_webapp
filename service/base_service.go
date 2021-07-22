package service

import "home/webonise/go/src/shopping_cart_app/model"

type Service struct {
	Accessor model.DBAccessor
}

type ServiceProvider interface {
	ProductServiceProvider
	ProductCartServiceProvider
}
