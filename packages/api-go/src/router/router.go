package router

import (
	"github.com/gorilla/mux"
	"github.com/otaviopontes/api-go/src/router/routes"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)

}
