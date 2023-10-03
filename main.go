package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {

	r := router.GenerateRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
