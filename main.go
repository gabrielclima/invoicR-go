package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDb()
	defer db.Close()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/invoices", RestInvoices).Methods("GET")
	router.HandleFunc("/invoices", RestCreateInvoice).Methods("POST")
	router.HandleFunc("/invoices/{document}", RestInvoiceByDoc).Methods("GET")
	// router.HandleFunc("/invoices/{document}", RestDeleteInvoice).Methods("DELETE")

	// colocando a API online
	log.Fatal(http.ListenAndServe(":8080", router))
}
