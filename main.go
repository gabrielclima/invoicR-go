package main

import (
	"log"
	http "net/http"
	"github.com/gabrielclima/go_rest_api/database"
	rest "github.com/gabrielclima/go_rest_api/rest"
)

func main() {
	InitDb()

	log.Fatal(http.ListenAndServe(":8080", Handlers()))
}
