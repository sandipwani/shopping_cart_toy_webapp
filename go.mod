module home/webonise/go/src/shopping_cart_app

go 1.16

replace (
	home/webonise/go/src/shopping_cart_app/containers => ./containers
	home/webonise/go/src/shopping_cart_app/controller => ./controller
	home/webonise/go/src/shopping_cart_app/database => ./database
	home/webonise/go/src/shopping_cart_app/migrations => ./migrations
	home/webonise/go/src/shopping_cart_app/model => ./model
	home/webonise/go/src/shopping_cart_app/router => ./router
	home/webonise/go/src/shopping_cart_app/service => ./service
	home/webonise/go/src/shopping_cart_app/structure => ./structure
	home/webonise/go/src/shopping_cart_app/view => ./view
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.2
	home/webonise/go/src/shopping_cart_app/controller v0.0.0-00010101000000-000000000000 // indirect
	home/webonise/go/src/shopping_cart_app/database v0.0.0-00010101000000-000000000000
	home/webonise/go/src/shopping_cart_app/model v0.0.0-00010101000000-000000000000 // indirect
	home/webonise/go/src/shopping_cart_app/router v0.0.0-00010101000000-000000000000
	home/webonise/go/src/shopping_cart_app/service v0.0.0-00010101000000-000000000000 // indirect
	home/webonise/go/src/shopping_cart_app/structure v0.0.0-00010101000000-000000000000 // indirect
)
