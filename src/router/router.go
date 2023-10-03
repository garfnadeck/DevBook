package router

import (
	"api/src/route"
	"github.com/gorilla/mux"
)

func GenerateRoutes() *mux.Router {
	r := mux.NewRouter()
	return route.ConfigRouter(r)

}
