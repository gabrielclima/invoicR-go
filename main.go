package main

import (
	"log"
	"net/http"
	"github.com/gabriellima/go_rest_api/database"
)

func main() {
	initDb()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", Handlers()))
}
