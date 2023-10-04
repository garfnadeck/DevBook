package main

import (
	"api/src/config"
	"api/src/router"
	"log"
	"net/http"
	"os"
)

func main() {
	config.LoadDatabase()
	r := router.GenerateRoutes()
	log.Fatal(http.ListenAndServe(os.Getenv("MAIN_PORT"), r))
}
