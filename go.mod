module home/webonise/go/src/shopping_cart_app

go 1.16

replace (
	home/webonise/go/src/shopping_cart_app/Database => ./Database
	home/webonise/go/src/shopping_cart_app/Module => ./Module
)

require (
	github.com/gorilla/mux v1.8.0
	home/webonise/go/src/shopping_cart_app/Database v0.0.0-00010101000000-000000000000 // indirect
	home/webonise/go/src/shopping_cart_app/Module v0.0.0-00010101000000-000000000000
)
