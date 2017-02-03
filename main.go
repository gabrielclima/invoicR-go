package main

import (
	_ "github.com/gabrielclima/go_rest_api/database"
	"github.com/gabrielclima/go_rest_api/rest"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", rest.Handlers()))
}
