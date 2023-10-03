package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

type RouteStructure struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
	return r
}
