package containers

import (
	"database/sql"
	"home/webonise/go/src/shopping_cart_app/controller"
	"home/webonise/go/src/shopping_cart_app/model"
	"home/webonise/go/src/shopping_cart_app/router"
	"home/webonise/go/src/shopping_cart_app/service"

	"github.com/gorilla/mux"
)

type Server struct {
	Router   *mux.Router
	MasterDB *sql.DB
}

func (s *Server) InitializeServer() {
	d := &model.WriteDbAccess{
		Db: s.MasterDB,
	}

	serv := &service.Service{
		Accessor: d,
	}

	sr := &router.SrvRouter{
		Router: s.Router,
		Controller: &controller.Srv{
			Service: serv,
		},
	}

	sr.InitilizeRoutes()
}
