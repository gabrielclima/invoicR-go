package main

import (
	"log"
	"net/http"
)

func main() {
	initDb()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", Handlers()))
}
