package main

import (
	"TesteWeb/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	api.Routes(router)

	log.Fatal(http.ListenAndServe(":7000", router))
}

