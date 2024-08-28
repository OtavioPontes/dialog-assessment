package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/otaviopontes/api-go/docs"
	middlewares "github.com/otaviopontes/api-go/src/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	Uri                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {

	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, routesPosts...)

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {

			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r

}
